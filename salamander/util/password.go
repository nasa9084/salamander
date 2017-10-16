package util

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

const (
	stretchN = 30
	hashAlg  = "sha512"
)

// Password hashes raw password string with salt(id)
func Password(password, id string) string {
	hashed := password
	for i := 0; i < stretchN; i++ {
		hashed = sha512HexDigest(hashed + id)
	}
	return fmt.Sprintf("$%s$%s$%d$%s", id, hashAlg, stretchN, hashed)
}

func sha512HexDigest(raw string) string {
	hash := sha512.New()
	hash.Write([]byte(raw))
	return hex.EncodeToString(hash.Sum(nil))
}
