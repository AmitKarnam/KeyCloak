package storagebackendvalidator

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/AmitKarnam/KeyCloak/internal/utils/logger/zapLogger"
	"github.com/AmitKarnam/KeyCloak/internal/utils/validators"
	"github.com/AmitKarnam/KeyCloak/models"
)

type StorageBackendValidator struct{}

var _ validators.Validator[models.StorageBackend] = &StorageBackendValidator{}

// Validate the data based on the type of storage backend type. Contains concrete implementations for each type of storage backend
func (sbv *StorageBackendValidator) Validate(data models.StorageBackend) error {

	// models.StorageBackend is of type json.RawMessage, needs to be marshalled to JSON so that type validators can be applied.
	// Once type checking happens, we need to hand off the storage backend related items to be validated for it's particular type
	// Thinking helper function will take map[string]interface{} as input and return error based on the validation

	zapLogger.KeyCloaklogger.Info("validating storage backend")
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	jsonData["storage_type"] = strings.ToLower((jsonData["storage_type"]).(string))

	switch jsonData["storage_type"] {
	case "file":
		err := fileValidator(jsonData)
		if err != nil {
			zapLogger.KeyCloaklogger.Errorf("%v : %v", "invalid file type detected", err)
			return err
		}
	default:
		return errors.New("storage backend type not supported")
	}

	return nil
}

// helper function to validate file type
func fileValidator(data map[string]interface{}) error {

	if data["file_path"].(string) == "" {
		return errors.New("empty file path not allowed")
	}

	return nil
}
