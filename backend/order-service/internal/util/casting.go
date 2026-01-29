package util

import "github.com/google/uuid"

func StringToUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

func UUIDToString(id uuid.UUID) string {
	return id.String()
}
