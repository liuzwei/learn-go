package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Md5Encrypt(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))

	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha256(password string) string {
	shaer := sha256.New()
	shaer.Write([]byte(password))
	return hex.EncodeToString(shaer.Sum(nil))
}

func Sha512(password string) string {
	shaer := sha512.New()
	shaer.Write([]byte(password))
	return hex.EncodeToString(shaer.Sum(nil))
}

func Sha1(password string) string {
	shaer := sha1.New()
	shaer.Write([]byte(password))
	return hex.EncodeToString(shaer.Sum(nil))
}
