package helper

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateRandomFilename() string {
	randomBytes := make([]byte, 8)
	_, _ = rand.Read(randomBytes)

	randomPart := hex.EncodeToString(randomBytes)
	timestamp := time.Now().UnixNano()

	return fmt.Sprintf("product-%d-%s.jpg", timestamp, randomPart)
}
