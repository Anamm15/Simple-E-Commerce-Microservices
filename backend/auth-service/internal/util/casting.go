package util

import "github.com/google/uuid"

func StringToUUID(uuidString string) (uuid.UUID, error) {
	return uuid.Parse(uuidString)
}
