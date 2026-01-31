package util

import (
	"crypto/rand"
	"encoding/base64"
	"io"
)

func GenerateRandomString(n int) (string, error) {
	b := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}
