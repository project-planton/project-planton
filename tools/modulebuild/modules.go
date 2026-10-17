package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// Engine is the IaC engine a module is written for. Its value is the module
// directory's name under iac/.
type Engine string

const (
	Pulumi    Engine = "pulumi"
	Terraform Engine = "tf"
)

func parseEngine(s string) (Engine, error) {
	switch s {
	case "pulumi":
		return Pulumi, nil
	case "terraform", "tf":
		return Terraform, nil
	}
	return "", fmt.Errorf("engine %q is neither pulumi nor terraform", s)
}

// Module is one official module a release ships.
type Module struct {
	// Dir is the module directory as the catalog names it, e.g.
	// catalog/aws/awsvpc/iac/tf -- the module manifest's key.
	Dir      string
	Provider string
	Kind     string
	Engine   Engine
}

// licenseFiles accompany every redistributed module at its archive root
// (Apache-2.0 section 4(d)).
var licenseFiles = []string{"LICENSE", "NOTICE"}

// ReleasedModules lists every module a release ships: each git-tracked
// catalog/<provider>/<kind>/iac/{pulumi,tf} directory. Providers whose name
// starts with an underscore (catalog/_test, catalog/_docs, ...) are internal
// and never released. Only tracked files count, so a build or plan left in a
// working tree never adds or changes a module.
func ReleasedModules(root string) ([]Module, error) {
	files, err := trackedFiles(root, "catalog")
	if err != nil {
		return nil, err
	}
	seen := map[string]Module{}
	for _, f := range files {
		parts := strings.Split(f, "/")
		if len(parts) < 6 || parts[3] != "iac" || strings.HasPrefix(parts[1], "_") {
			continue
		}
		engine := Engine(parts[4])
		if engine != Pulumi && engine != Terraform {
			continue
		}
		dir := strings.Join(parts[:5], "/")
		seen[dir] = Module{Dir: dir, Provider: parts[1], Kind: parts[2], Engine: engine}
	}
	modules := make([]Module, 0, len(seen))
	for _, m := range seen {
		modules = append(modules, m)
	}
	sort.Slice(modules, func(i, j int) bool { return modules[i].Dir < modules[j].Dir })
	return modules, nil
}

// executedFiles is what a module archive a runner executes ships: the
// module's tracked files except documentation, plus the license files.
// Markdown is left out so editing a module's README never changes the bytes
// a runner executes, and so never moves the module's digest.
func executedFiles(root, dir string) (map[string][]byte, error) {
	return moduleFiles(root, dir, func(rel string) bool {
		return !strings.EqualFold(filepath.Ext(rel), ".md")
	})
}

// sourceFiles is what a module's source archive ships: every tracked file of
// the module, documentation included, plus the license files. Nothing
// executes it, so it plays no part in the module's digest.
func sourceFiles(root, dir string) (map[string][]byte, error) {
	return moduleFiles(root, dir, func(string) bool { return true })
}

func moduleFiles(root, dir string, keep func(rel string) bool) (map[string][]byte, error) {
	dir = strings.TrimSuffix(filepath.ToSlash(dir), "/")
	tracked, err := trackedFiles(root, dir)
	if err != nil {
		return nil, err
	}
	if len(tracked) == 0 {
		return nil, fmt.Errorf("%s holds no tracked files", dir)
	}
	entries := map[string][]byte{}
	for _, f := range tracked {
		rel := strings.TrimPrefix(f, dir+"/")
		if !keep(rel) {
			continue
		}
		content, err := os.ReadFile(filepath.Join(root, f))
		if err != nil {
			return nil, err
		}
		entries[rel] = content
	}
	for _, name := range licenseFiles {
		if _, clash := entries[name]; clash {
			return nil, fmt.Errorf("%s has its own %s, which would hide the repository's", dir, name)
		}
		content, err := os.ReadFile(filepath.Join(root, name))
		if err != nil {
			return nil, err
		}
		entries[name] = content
	}
	return entries, nil
}

// trackedFiles lists the git-tracked files under path, repository-relative.
func trackedFiles(root, path string) ([]string, error) {
	cmd := exec.Command("git", "ls-files", "-z", "--", path)
	cmd.Dir = root
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git ls-files %s: %v: %s", path, err, stderr.String())
	}
	var files []string
	for _, f := range strings.Split(string(out), "\x00") {
		if f != "" {
			files = append(files, f)
		}
	}
	return files, nil
}
