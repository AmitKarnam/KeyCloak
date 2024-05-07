package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type EncryptionStrategy string
type StorageBackend json.RawMessage

const (
	AES EncryptionStrategy = "aes"
)

type SecretEngine struct {
	gorm.Model
	Name                string             `json:"name"`
	Encryption_Strategy EncryptionStrategy `json:"encryption_strategy"`
	Storage_Backend     StorageBackend     `json:"storage_backend"`
}
