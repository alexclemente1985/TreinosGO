package id

import (
	"crypto/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	entropy = ulid.Monotonic(rand.Reader, 0)
	mu      sync.Mutex
)

// New gera um ULID seguro para concorrência
func New() string {
	mu.Lock()
	defer mu.Unlock()

	return ulid.MustNew(
		ulid.Timestamp(time.Now().UTC()),
		entropy,
	).String()
}
