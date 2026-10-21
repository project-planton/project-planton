// Command modulebuild is the one place a release builds, packs and describes
// Planton's platform IaC modules.
//
// A platform recognises a module by its digest: the checksums of the artifacts
// a release published for it (see pkg/downloads/modulemanifest). For that to
// say "unchanged" whenever a module's source did not change, every artifact
// must be a pure function of its source -- so this tool, and nothing else in
// the release, decides which files a module ships, how a Pulumi binary is
// compiled and compressed, and how an archive is written. The release
// workflows call it and only upload what it produces.
//
// Subcommands:
//
//	list      -engine pulumi|terraform [-provider p]   the module directories a release ships
//	pulumi    -module <dir> -platform <goos_goarch> -out <file.gz>
//	terraform -module <dir> -out <module.zip>
//	source    -module <dir> -out <source.zip>
//	manifest  -engine pulumi|terraform -tag <tag> -parts <dir> -out <manifest.json>
//
// Every subcommand runs from the repository root.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/plantonhq/planton/pkg/downloads/modulemanifest"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	var err error
	switch os.Args[1] {
	case "list":
		err = runList(os.Args[2:])
	case "pulumi":
		err = runPulumi(os.Args[2:])
	case "terraform":
		err = runTerraform(os.Args[2:])
	case "source":
		err = runSource(os.Args[2:])
	case "manifest":
		err = runManifest(os.Args[2:])
	default:
		usage()
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "modulebuild %s: %v\n", os.Args[1], err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: modulebuild list|pulumi|terraform|source|manifest [flags]")
	os.Exit(2)
}

func runList(args []string) error {
	fs := flag.NewFlagSet("list", flag.ExitOnError)
	engine := fs.String("engine", "", "pulumi or terraform")
	provider := fs.String("provider", "", "only this provider's modules")
	_ = fs.Parse(args)
	e, err := modulemanifest.ParseEngine(*engine)
	if err != nil {
		return err
	}
	modules, err := ReleasedModules(".")
	if err != nil {
		return err
	}
	for _, m := range modules {
		if m.Engine == e && (*provider == "" || m.Provider == *provider) {
			fmt.Println(m.Dir)
		}
	}
	return nil
}

func runPulumi(args []string) error {
	fs := flag.NewFlagSet("pulumi", flag.ExitOnError)
	module := fs.String("module", "", "module directory, e.g. catalog/aws/awsvpc/iac/pulumi")
	platform := fs.String("platform", "", "target platform, e.g. linux_amd64")
	out := fs.String("out", "", "gzipped binary to write")
	_ = fs.Parse(args)
	if *module == "" || *platform == "" || *out == "" {
		return fmt.Errorf("-module, -platform and -out are required")
	}
	goos, goarch, ok := strings.Cut(*platform, "_")
	if !ok {
		return fmt.Errorf("platform %q is not <goos>_<goarch>", *platform)
	}
	return BuildPulumiBinary(".", *module, goos, goarch, *out)
}

func runTerraform(args []string) error {
	fs := flag.NewFlagSet("terraform", flag.ExitOnError)
	module := fs.String("module", "", "module directory, e.g. catalog/aws/awsvpc/iac/tf")
	out := fs.String("out", "", "module.zip to write")
	_ = fs.Parse(args)
	if *module == "" || *out == "" {
		return fmt.Errorf("-module and -out are required")
	}
	return PackModule(".", *module, *out)
}

func runSource(args []string) error {
	fs := flag.NewFlagSet("source", flag.ExitOnError)
	module := fs.String("module", "", "module directory, e.g. catalog/aws/awsvpc/iac/pulumi")
	out := fs.String("out", "", "source.zip to write")
	_ = fs.Parse(args)
	if *module == "" || *out == "" {
		return fmt.Errorf("-module and -out are required")
	}
	return PackSource(".", *module, *out)
}

func runManifest(args []string) error {
	fs := flag.NewFlagSet("manifest", flag.ExitOnError)
	engine := fs.String("engine", "", "pulumi or terraform: whose modules the manifest describes")
	tag := fs.String("tag", "", "release tag the manifest describes")
	parts := fs.String("parts", "", "directory holding every leg's published-artifact list for the engine")
	out := fs.String("out", "", "manifest.json to write")
	_ = fs.Parse(args)
	if *engine == "" || *tag == "" || *parts == "" || *out == "" {
		return fmt.Errorf("-engine, -tag, -parts and -out are required")
	}
	e, err := modulemanifest.ParseEngine(*engine)
	if err != nil {
		return err
	}
	modules, err := ReleasedModules(".")
	if err != nil {
		return err
	}
	published, err := ReadParts(*parts)
	if err != nil {
		return err
	}
	data, err := BuildManifest(*tag, e, modules, published)
	if err != nil {
		return err
	}
	return os.WriteFile(*out, data, 0o644)
}
