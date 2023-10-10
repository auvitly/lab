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

// assistantKey - key for the context in which the assistant will be stored.
var assistantKey = newKey()
