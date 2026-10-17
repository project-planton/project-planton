package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// buildFlags compile a Pulumi module into a binary that depends only on its
// source: no build paths (-trimpath), no VCS stamp (-buildvcs=false), no
// build id (-buildid=), no symbol or debug tables (-s -w).
var buildFlags = []string{"-trimpath", "-buildvcs=false", "-ldflags=-s -w -buildid="}

// BuildPulumiBinary compiles a Pulumi module for goos/goarch and writes it
// gzipped to out.
//
// Everything that could make two builds of one source differ is fixed: the
// toolchain is the one go.mod names (and this tool must run on it, so its own
// gzip writes the same bytes), the workspace file is ignored, cgo is off, the
// microarchitecture levels are the defaults, and no ambient GOFLAGS or
// experiments reach the compiler.
func BuildPulumiBinary(root, dir, goos, goarch, out string) error {
	toolchain, err := goModToolchain(root)
	if err != nil {
		return err
	}
	if runtime.Version() != toolchain {
		return fmt.Errorf("modulebuild runs on %s but go.mod names %s; run it with GOTOOLCHAIN=%s so its gzip output is the release's", runtime.Version(), toolchain, toolchain)
	}
	tmp, err := os.MkdirTemp("", "modulebuild-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	binary := filepath.Join(tmp, "module")

	args := append([]string{"build"}, buildFlags...)
	args = append(args, "-o", binary, "./"+strings.TrimPrefix(filepath.ToSlash(dir), "./"))
	cmd := exec.Command("go", args...)
	cmd.Dir = root
	cmd.Env = buildEnv(os.Environ(), toolchain, goos, goarch)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("building %s for %s_%s: %w", dir, goos, goarch, err)
	}
	content, err := os.ReadFile(binary)
	if err != nil {
		return err
	}
	compressed, err := gzipBytes(content)
	if err != nil {
		return err
	}
	return os.WriteFile(out, compressed, 0o644)
}

// buildEnv is the ambient environment with every build-affecting variable
// replaced by the release's fixed value.
func buildEnv(ambient []string, toolchain, goos, goarch string) []string {
	fixed := map[string]string{
		"CGO_ENABLED":  "0",
		"GOOS":         goos,
		"GOARCH":       goarch,
		"GOTOOLCHAIN":  toolchain,
		"GOWORK":       "off",
		"GOFLAGS":      "-mod=readonly",
		"GOEXPERIMENT": "",
		"GOAMD64":      "v1",
		"GOARM64":      "v8.0",
	}
	env := make([]string, 0, len(ambient)+len(fixed))
	for _, kv := range ambient {
		name, _, _ := strings.Cut(kv, "=")
		if _, overridden := fixed[name]; !overridden {
			env = append(env, kv)
		}
	}
	for name, value := range fixed {
		env = append(env, name+"="+value)
	}
	return env
}

// gzipBytes compresses at the best level with an empty header: no file name,
// no modification time, so the output is a pure function of the input for a
// given Go release.
func gzipBytes(content []byte) ([]byte, error) {
	var buf bytes.Buffer
	zw, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return nil, err
	}
	if _, err := zw.Write(content); err != nil {
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// goModToolchain is the exact toolchain go.mod names: its toolchain line, or
// else its go line.
func goModToolchain(root string) (string, error) {
	f, err := os.Open(filepath.Join(root, "go.mod"))
	if err != nil {
		return "", err
	}
	defer f.Close()
	var goLine, toolchainLine string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			continue
		}
		switch fields[0] {
		case "go":
			goLine = "go" + fields[1]
		case "toolchain":
			toolchainLine = fields[1]
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	if toolchainLine != "" {
		return toolchainLine, nil
	}
	if goLine == "" {
		return "", fmt.Errorf("go.mod names no go version")
	}
	return goLine, nil
}
