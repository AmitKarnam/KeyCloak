package storagebackendvalidator

import (
	"github.com/AmitKarnam/KeyCloak/internal/utils/validators"
	"github.com/AmitKarnam/KeyCloak/models"
)

type StorageBackendValidator struct{}

var _ validators.Validator[models.StorageBackend] = &StorageBackendValidator{}

// Validate the data based on the type of storage backend type. Contains concrete implementations for each type of storage backend
func (sbv *StorageBackendValidator) Validate(data models.StorageBackend) error {

	return nil
}

func fileValidator(data models.StorageBackend) error { return nil }
