package detzip

import (
	"archive/zip"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"testing"
)

var sample = map[string][]byte{
	"main.tf":            []byte("resource \"null_resource\" \"x\" {}\n"),
	"LICENSE":            []byte("license\n"),
	"hack/manifest.yaml": []byte("kind: Manifest\n"),
}

// TestStoreGoldenBytes pins a stored archive's exact bytes: they depend on
// nothing but the entries, not the Go release that wrote them, so a change
// here is a change in every published module archive.
func TestStoreGoldenBytes(t *testing.T) {
	archive, err := Bytes(sample, Store)
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(archive)
	if got, want := hex.EncodeToString(sum[:]), "b6177a4a94a5a2551e1d98bfd2935443ae70117db4d304fdee382e7bf2149886"; got != want {
		t.Fatalf("stored archive sha256 = %s, want %s", got, want)
	}
}

func TestSameEntriesSameBytes(t *testing.T) {
	for _, method := range []Method{Store, Deflate} {
		first, err := Bytes(sample, method)
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 10; i++ {
			// A fresh map each time: Go randomizes map iteration, so an
			// order-dependent writer would show up here.
			again := map[string][]byte{}
			for k, v := range sample {
				again[k] = v
			}
			got, err := Bytes(again, method)
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, first) {
				t.Fatalf("method %d: the same entries packed to different bytes", method)
			}
		}
	}
}

func TestArchiveReadsBackInOrderWithNoTimestamps(t *testing.T) {
	for _, method := range []Method{Store, Deflate} {
		archive, err := Bytes(sample, method)
		if err != nil {
			t.Fatal(err)
		}
		r, err := zip.NewReader(bytes.NewReader(archive), int64(len(archive)))
		if err != nil {
			t.Fatal(err)
		}
		var names []string
		for _, f := range r.File {
			names = append(names, f.Name)
			if f.ModifiedDate != 0 || f.ModifiedTime != 0 || len(f.Extra) != 0 {
				t.Errorf("%s: carries a timestamp or extra field; a rebuild would not reproduce it", f.Name)
			}
			rc, err := f.Open()
			if err != nil {
				t.Fatal(err)
			}
			content, err := io.ReadAll(rc)
			rc.Close()
			if err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(content, sample[f.Name]) {
				t.Errorf("%s: content does not round-trip", f.Name)
			}
		}
		want := []string{"LICENSE", "hack/manifest.yaml", "main.tf"}
		if len(names) != len(want) {
			t.Fatalf("entries %v, want %v", names, want)
		}
		for i := range want {
			if names[i] != want[i] {
				t.Fatalf("entries %v, want lexical order %v", names, want)
			}
		}
	}
}
