package utils

import "github.com/google/uuid"

func CreateNewUUID() string {
	return uuid.New().String()
}
