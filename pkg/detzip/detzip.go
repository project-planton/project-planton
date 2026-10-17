// Package detzip writes zip archives whose bytes are a pure function of their
// entries' names and contents.
//
// Every archive a release publishes under a checksum is written here: module
// archives, definitions archives and the catalog bundle. A rebuild of unchanged
// content must reproduce the published bytes exactly -- that is what lets a
// release say "this module did not change" by comparing checksums, and what
// lets a consumer re-pack and verify an archive. So nothing that varies between
// two builds of the same tree reaches the archive: entries are written in
// lexical order, every entry carries the format's zero timestamp and no extra
// fields, no owner, group or permission bits are recorded, and each entry is
// written raw with its sizes and CRC known up front (descriptor-free entries
// are also the most widely readable ZIP shape).
package detzip

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"fmt"
	"hash/crc32"
	"io"
	"sort"
)

// Method is how an archive stores its entries.
type Method uint16

const (
	// Store keeps entries uncompressed. Its bytes depend on nothing but the
	// content -- not even the Go release that wrote them -- so it is the
	// method for archives whose checksum must survive toolchain upgrades.
	Store Method = Method(zip.Store)
	// Deflate compresses entries. Its bytes are reproducible for a given Go
	// release; use it where size matters more than checksum stability across
	// toolchains.
	Deflate Method = Method(zip.Deflate)
)

// Write packs entries (archive path -> content) into w.
func Write(w io.Writer, entries map[string][]byte, method Method) error {
	names := make([]string, 0, len(entries))
	for name := range entries {
		names = append(names, name)
	}
	sort.Strings(names)

	zw := zip.NewWriter(w)
	for _, name := range names {
		content := entries[name]
		stored, err := encode(content, method)
		if err != nil {
			return fmt.Errorf("compressing %s: %w", name, err)
		}
		entry, err := zw.CreateRaw(&zip.FileHeader{
			Name:               name,
			Method:             uint16(method),
			CRC32:              crc32.ChecksumIEEE(content),
			CompressedSize64:   uint64(len(stored)),
			UncompressedSize64: uint64(len(content)),
		})
		if err != nil {
			return fmt.Errorf("creating archive entry %s: %w", name, err)
		}
		if _, err := entry.Write(stored); err != nil {
			return fmt.Errorf("writing archive entry %s: %w", name, err)
		}
	}
	return zw.Close()
}

// Bytes packs entries into an in-memory archive.
func Bytes(entries map[string][]byte, method Method) ([]byte, error) {
	var buf bytes.Buffer
	if err := Write(&buf, entries, method); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// encode returns the entry's bytes as the archive stores them.
func encode(content []byte, method Method) ([]byte, error) {
	switch method {
	case Store:
		return content, nil
	case Deflate:
		var buf bytes.Buffer
		fw, err := flate.NewWriter(&buf, flate.DefaultCompression)
		if err != nil {
			return nil, err
		}
		if _, err := fw.Write(content); err != nil {
			return nil, err
		}
		if err := fw.Close(); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	default:
		return nil, fmt.Errorf("unsupported zip method %d", method)
	}
}
