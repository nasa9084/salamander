package util

import (
	"crypto/sha512"
	"encoding/hex"
	"strconv"
)

const (
	stretchN = 30
	hashAlg = "sha512"
)

// Password hashes raw password string with salt(id)
func Password(password, id string) string {
	hashed := password
	for i := 0; i < stretchN; i++ {
		hashed = sha512HexDigest(hashed + id)
	}
	return "$" + id + "$" + hashAlg + "$" + strconv.Itoa(stretchN) + "$" + hashed
}

func sha512HexDigest(raw string) string {
	hash := sha512.New()
	hash.Write([]byte(raw))
	return hex.EncodeToString(hash.Sum(nil))
}
