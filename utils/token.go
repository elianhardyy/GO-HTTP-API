package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateTokenEmail(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}