package gopeach

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

const (
	// DefaultSaltLength is the default salt length.
	DefaultSaltLength = 32
	// DefaultHashLength is the default hash length.
	DefaultHashLength = 64
)

// Hash hashes a string.
func Hash(s string) (string, error) {
	salt := make([]byte, DefaultSaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash, err := scrypt.Key([]byte(s), salt, 1024*DefaultHashLength, 8, 1, DefaultHashLength)
	return base64.StdEncoding.EncodeToString(salt) + ":" + base64.StdEncoding.EncodeToString(hash), err // "+"" is 99x faster than "fmt.Sprintf"
}

// Compare compares a string and a hash.
func CompareHash(s, hash string) (bool, error) {
	var err error
	hparts := strings.Split(hash, ":")
	if len(hparts) != 2 {
		return false, fmt.Errorf("invalid hash")
	}
	salt, errs := base64.StdEncoding.DecodeString(hparts[0])
	h, errh := base64.StdEncoding.DecodeString(hparts[1])
	if errs != nil || errh != nil {
		err = errs
		if err == nil {
			err = errh
		}
		return false, err
	}
	hashlen := len(h)
	h2, err := scrypt.Key([]byte(s), salt, 1024*hashlen, 8, 1, hashlen)
	if err != nil {
		return false, err
	}

	return bytes.Compare(h, h2) == 0, nil
}
