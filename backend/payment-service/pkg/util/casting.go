package util

import "github.com/google/uuid"

func StringToUUID(uuidStr string) (uuid.UUID, error) {
	return uuid.Parse(uuidStr)
}
