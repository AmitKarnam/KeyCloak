package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type EncryptionStrategy string

type StorageBackend json.RawMessage

func (sb StorageBackend) MarshalJSON() ([]byte, error) {
	return []byte(sb), nil
}
func (sb *StorageBackend) UnmarshalJSON(data []byte) error {
	*sb = append((*sb)[0:0], data...)
	return nil
}

const (
	AES EncryptionStrategy = "aes"
)

type SecretEngine struct {
	gorm.Model
	Name                string             `json:"name"`
	Encryption_Strategy EncryptionStrategy `json:"encryption_strategy"`
	Storage_Backend     StorageBackend     `json:"storage_backend"`
}
