// Command providerlocks keeps the catalog's OpenTofu provider locks current.
// It is the thin operational face of pkg/iac/tofu/providerlocks:
//
//	go run ./pkg/iac/tofu/providerlocks/cli lock  [--upgrade aws,helm]   (make lock-providers)
//	go run ./pkg/iac/tofu/providerlocks/cli check                     (hack/guards/ensure_tf_provider_locks.sh)
//
// lock keeps every locked version unless a module's constraint excludes it or
// --upgrade names the provider; only those providers reach the registry and
// `tofu providers lock`. It then rewrites every module's lock as a projection
// of the canonical one. check is offline and writes nothing.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/plantonhq/planton/pkg/iac/tofu/providerlocks"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: providerlocks lock|check [flags]")
	}
	switch os.Args[1] {
	case "lock":
		fs := flag.NewFlagSet("lock", flag.ExitOnError)
		repoRoot := fs.String("repo-root", ".", "repository root")
		upgrade := fs.String("upgrade", "", "comma-separated providers to move to their newest allowed release (\"aws\" or an address)")
		scratch := fs.String("scratch", "", "directory for the throwaway modules `tofu providers lock` runs in (default: a temp dir)")
		if err := fs.Parse(os.Args[2:]); err != nil {
			return err
		}
		return lock(*repoRoot, splitList(*upgrade), *scratch)
	case "check":
		fs := flag.NewFlagSet("check", flag.ExitOnError)
		repoRoot := fs.String("repo-root", ".", "repository root")
		if err := fs.Parse(os.Args[2:]); err != nil {
			return err
		}
		problems, err := providerlocks.Check(*repoRoot)
		if err != nil {
			return err
		}
		if len(problems) > 0 {
			for _, problem := range problems {
				fmt.Fprintln(os.Stderr, "  - "+problem)
			}
			return fmt.Errorf("%d provider lock problem(s)", len(problems))
		}
		fmt.Println("every OpenTofu module's provider lock is current")
		return nil
	default:
		return fmt.Errorf("unknown command %q: want lock or check", os.Args[1])
	}
}

func lock(repoRoot string, upgrade []string, scratch string) error {
	modules, err := providerlocks.Modules(repoRoot)
	if err != nil {
		return err
	}
	canonicalPath := filepath.Join(repoRoot, providerlocks.CanonicalLockPath)
	canonical, err := providerlocks.ReadLock(canonicalPath)
	if err != nil {
		return err
	}
	moves, err := providerlocks.PlanMoves(modules, canonical, upgrade, providerlocks.HTTPRegistry{})
	if err != nil {
		return err
	}
	if len(moves) > 0 {
		if scratch == "" {
			if scratch, err = os.MkdirTemp("", "providerlocks-"); err != nil {
				return err
			}
			defer os.RemoveAll(scratch)
		}
		for _, move := range moves {
			from := move.From
			if from == "" {
				from = "(none)"
			}
			fmt.Printf("%s: %s -> %s (%s; %s)\n", move.Address, from, move.Version, move.Reason, strings.Join(move.Platforms, ", "))
			locked, err := providerlocks.RecordChecksums(move, scratch)
			if err != nil {
				return err
			}
			canonical[move.Address] = locked
		}
	}
	// A provider no module requires any more leaves the catalog.
	required := providerlocks.ConstraintsByAddress(modules)
	for address := range canonical {
		if _, ok := required[address]; !ok {
			fmt.Printf("%s: no module requires it; removed\n", address)
			delete(canonical, address)
		}
	}
	if err := os.MkdirAll(filepath.Dir(canonicalPath), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(canonicalPath, providerlocks.Render(canonical), 0o644); err != nil {
		return err
	}
	if err := providerlocks.WriteProjections(repoRoot, canonical, modules); err != nil {
		return err
	}
	fmt.Printf("locked %d providers for %d modules\n", len(canonical), len(modules))
	return nil
}

func splitList(raw string) []string {
	var out []string
	for _, item := range strings.Split(raw, ",") {
		if item = strings.TrimSpace(item); item != "" {
			out = append(out, item)
		}
	}
	return out
}
