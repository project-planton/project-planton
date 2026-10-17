package main

import (
	"os"

	"github.com/plantonhq/planton/pkg/detzip"
)

// PackModule writes the archive a runner executes for an OpenTofu/Terraform
// module (module.zip). Entries are stored, not deflated, so the bytes do not
// move with the Go release that writes them.
func PackModule(root, dir, out string) error {
	entries, err := executedFiles(root, dir)
	if err != nil {
		return err
	}
	return writeArchive(entries, out)
}

// PackSource writes a module's source archive (source.zip), the module's
// tracked source tree for ejecting and reading.
func PackSource(root, dir, out string) error {
	entries, err := sourceFiles(root, dir)
	if err != nil {
		return err
	}
	return writeArchive(entries, out)
}

func writeArchive(entries map[string][]byte, out string) error {
	f, err := os.Create(out)
	if err != nil {
		return err
	}
	if err := detzip.Write(f, entries, detzip.Store); err != nil {
		f.Close()
		return err
	}
	return f.Close()
}
