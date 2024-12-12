package ulidgen

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

// ULIDGenerator holds the state necessary for generating ULIDs.
type ULIDGenerator struct {
	entropy *ulid.MonotonicEntropy
}

// NewGenerator initializes a new ULIDGenerator with its entropy source.
func NewGenerator() *ULIDGenerator {
	t := time.Now().UTC()
	source := rand.New(rand.NewSource(t.UnixNano()))
	entropy := ulid.Monotonic(source, 0)

	return &ULIDGenerator{
		entropy: entropy,
	}
}

// Generate creates a new ULID using the generator's entropy source.
func (g *ULIDGenerator) Generate() ulid.ULID {
	t := time.Now().UTC()
	return ulid.MustNew(ulid.Timestamp(t), g.entropy)
}
