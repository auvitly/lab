package assistant

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func newKey() string {
	var sha = sha256.New()

	sha.Write([]byte(time.Now().String()))

	return hex.EncodeToString(sha.Sum([]byte{}))
}

var (
	// Key for the context in which the assistant will be stored.
	// ! Should not be changed.
	key = newKey()
)
