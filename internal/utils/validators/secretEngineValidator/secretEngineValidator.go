package secretenginevalidator

import (
	"errors"

	"github.com/AmitKarnam/KeyCloak/internal/utils/validators"
	storagebackendvalidator "github.com/AmitKarnam/KeyCloak/internal/utils/validators/storageBackendValidator"
	"github.com/AmitKarnam/KeyCloak/models"
)

type secretEngineValidator struct{}

var _ validators.Validator[models.SecretEngine] = &secretEngineValidator{}

// Validate if the secret engine model instance passed is valid or not.
func (sev *secretEngineValidator) Validate(data models.SecretEngine) error {

	// Need to check the validity of the Secret Engine Name
	if data.Name == "" {
		return errors.New("empty secret engine name value not allowed")
	}

	// Need to check the validity of the Secret Engine Encryption Strategy
	if data.Encryption_Strategy == "" {
		return errors.New("empty secret engine encryption strategy value not allowed")
	}

	// Need to check the validity of the Secret Engine Storage Backend ( Call the Storage Backend Validator )
	sbv := storagebackendvalidator.StorageBackendValidator{}
	err := sbv.Validate(data.Storage_Backend)
	if err != nil {
		return err
	}

	return nil
}
