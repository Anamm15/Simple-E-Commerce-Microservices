package util

import (
	"io"
	"bytes"

	"github.com/google/uuid"
)

func StringToUUID(uuidString string) (uuid.UUID, error) {
	return uuid.Parse(uuidString)
}

func UUIDToString(uuid uuid.UUID) string {
	return uuid.String()
}

func ByteToIOReader(b []byte) io.Reader {
	return bytes.NewReader(b)
}
