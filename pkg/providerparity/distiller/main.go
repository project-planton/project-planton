//go:build !codegen
// +build !codegen

// Distills Terraform provider schemas into pkg/providerparity's committed
// artifacts.
//
// The provider is the only authority on its own configurable surface, and it
// is self-describing: `providers schema -json` emits every resource, argument,
// type, and flag at an exact released version. Committing that output whole
// would carry descriptions and data sources the parity accounting never
// reads; this tool resolves each provider in a throwaway OpenTofu workspace,
// reduces the schema to the configurable surface, and writes one
// deterministic gzipped artifact per provider (see
// providerparity.WriteSchemaFile for the determinism and replace-on-repin
// contract).
//
// The version is never chosen here. Every OpenTofu module runs the exact
// release the catalog's canonical lock names
// (catalog/_providers/.terraform.lock.hcl), so parity is measured at that
// release and no other: the workspace pins each provider to its locked
// version, and a provider the lock does not carry is refused.
//
// Run via Makefile: make generate-provider-schemas
//
//	go run ./pkg/providerparity/distiller \
//	    --out-dir pkg/providerparity/schemas \
//	    --provider 'google=hashicorp/google' \
//	    --provider 'google-beta=hashicorp/google-beta'
//
// Moving a provider is a lock sweep (`make lock-providers upgrade=<p>`)
// followed by a re-run: the artifact for the same provider is replaced, never
// accumulated, and the resulting parity-check failures ARE the migration work
// list.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/plantonhq/planton/pkg/iac/tofu/providerlocks"
	"github.com/plantonhq/planton/pkg/providerparity"
)

// pin is one provider to distill: the local/artifact name, the registry
// source address, and the exact version the canonical lock names for it.
type pin struct {
	name    string
	source  string
	version string
}

type pinList []pin

func (p *pinList) String() string {
	var parts []string
	for _, x := range *p {
		parts = append(parts, fmt.Sprintf("%s=%s", x.name, x.source))
	}
	return strings.Join(parts, ",")
}

func (p *pinList) Set(v string) error {
	name, source, ok := strings.Cut(v, "=")
	if !ok || name == "" || source == "" || strings.Contains(source, "@") {
		return fmt.Errorf("provider %q: want name=source (the version comes from the canonical lock)", v)
	}
	*p = append(*p, pin{name: name, source: source})
	return nil
}

func main() {
	var outDir, lockPath string
	var pins pinList
	flag.StringVar(&outDir, "out-dir", providerparity.DefaultSchemaDir, "directory for the committed artifacts")
	flag.StringVar(&lockPath, "lock", providerlocks.CanonicalLockPath, "the canonical provider lock that names each provider's version")
	flag.Var(&pins, "provider", "provider to distill as name=source (repeatable)")
	flag.Parse()
	if len(pins) == 0 {
		fatal("at least one --provider is required")
	}

	lock, err := providerlocks.ReadLock(lockPath)
	if err != nil {
		fatal("reading %s: %v", lockPath, err)
	}
	for i, p := range pins {
		version, ok := providerparity.LockedVersion(lock, p.source)
		if !ok {
			fatal("%s locks no provider %q -- lock it with `make lock-providers` before distilling", lockPath, p.source)
		}
		pins[i].version = version
	}

	workDir, err := os.MkdirTemp("", "providerparity-distill-*")
	if err != nil {
		fatal("creating workspace: %v", err)
	}
	defer os.RemoveAll(workDir)

	if err := os.WriteFile(filepath.Join(workDir, "providers.tf"), []byte(workspaceConfig(pins)), 0o644); err != nil {
		fatal("writing workspace config: %v", err)
	}
	if out, err := tofu(workDir, "init", "-backend=false", "-input=false", "-no-color"); err != nil {
		fatal("tofu init: %v\n%s", err, out)
	}
	schemaJSON, err := tofu(workDir, "providers", "schema", "-json")
	if err != nil {
		fatal("tofu providers schema: %v\n%s", err, schemaJSON)
	}

	var raw rawSchemas
	if err := json.Unmarshal(schemaJSON, &raw); err != nil {
		fatal("parsing schema JSON: %v", err)
	}

	for _, p := range pins {
		body := raw.forSource(p.source)
		if body == nil {
			fatal("schema output carries no provider matching source %q", p.source)
		}
		schema := &providerparity.Schema{
			Provider:       p.name,
			Source:         p.source,
			Version:        p.version,
			ProviderConfig: distillBlock(body.Provider.Block),
			Resources:      map[string]*providerparity.Block{},
		}
		for name, rs := range body.ResourceSchemas {
			schema.Resources[name] = distillBlock(rs.Block)
		}
		path, err := providerparity.WriteSchemaFile(outDir, schema)
		if err != nil {
			fatal("writing artifact for %s: %v", p.name, err)
		}
		fmt.Printf("distilled %s %s: %d resources -> %s\n", p.name, p.version, len(schema.Resources), path)
	}
}

// workspaceConfig renders the throwaway workspace's required_providers block,
// each provider pinned to exactly its locked version.
func workspaceConfig(pins pinList) string {
	var b strings.Builder
	b.WriteString("terraform {\n  required_providers {\n")
	for _, p := range pins {
		fmt.Fprintf(&b, "    %s = {\n      source  = %q\n      version = %q\n    }\n", p.name, p.source, p.version)
	}
	b.WriteString("  }\n}\n")
	return b.String()
}

func tofu(dir string, args ...string) ([]byte, error) {
	cmd := exec.Command("tofu", args...)
	cmd.Dir = dir
	return cmd.CombinedOutput()
}

// rawSchemas mirrors the slice of `providers schema -json` output the
// distillation reads. Everything else (descriptions, data sources, functions)
// is deliberately not modeled.
type rawSchemas struct {
	ProviderSchemas map[string]*rawProviderBody `json:"provider_schemas"`
}

func (r rawSchemas) forSource(source string) *rawProviderBody {
	for key, body := range r.ProviderSchemas {
		if strings.HasSuffix(key, "/"+source) {
			return body
		}
	}
	return nil
}

type rawProviderBody struct {
	// Provider is the provider's own configuration block -- the arguments a
	// `provider "<name>" {}` block accepts. Part of the parity surface: the
	// provider-config accounting reads it exactly like resource blocks.
	Provider        rawResourceSchema            `json:"provider"`
	ResourceSchemas map[string]rawResourceSchema `json:"resource_schemas"`
}

type rawResourceSchema struct {
	Block rawBlock `json:"block"`
}

type rawBlock struct {
	Attributes map[string]rawAttribute `json:"attributes"`
	BlockTypes map[string]rawBlockType `json:"block_types"`
	Deprecated bool                    `json:"deprecated"`
}

type rawAttribute struct {
	Type       json.RawMessage `json:"type"`
	Required   bool            `json:"required"`
	Optional   bool            `json:"optional"`
	Computed   bool            `json:"computed"`
	Sensitive  bool            `json:"sensitive"`
	Deprecated bool            `json:"deprecated"`
}

type rawBlockType struct {
	NestingMode string   `json:"nesting_mode"`
	Block       rawBlock `json:"block"`
}

func distillBlock(rb rawBlock) *providerparity.Block {
	b := &providerparity.Block{Deprecated: rb.Deprecated}
	if len(rb.Attributes) > 0 {
		b.Attributes = map[string]*providerparity.Attribute{}
		for name, a := range rb.Attributes {
			b.Attributes[name] = &providerparity.Attribute{
				Type:       a.Type,
				Required:   a.Required,
				Optional:   a.Optional,
				Computed:   a.Computed,
				Sensitive:  a.Sensitive,
				Deprecated: a.Deprecated,
			}
		}
	}
	if len(rb.BlockTypes) > 0 {
		b.Blocks = map[string]*providerparity.NestedBlock{}
		for name, bt := range rb.BlockTypes {
			b.Blocks[name] = &providerparity.NestedBlock{
				NestingMode: bt.NestingMode,
				Block:       distillBlock(bt.Block),
			}
		}
	}
	return b
}

func fatal(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "providerparity distiller: "+format+"\n", args...)
	os.Exit(1)
}
