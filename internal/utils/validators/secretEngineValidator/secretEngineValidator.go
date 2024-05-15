package secretenginevalidator

import (
	"errors"

	"github.com/AmitKarnam/KeyCloak/internal/internalerrors"
	"github.com/AmitKarnam/KeyCloak/internal/utils/logger/zapLogger"
	"github.com/AmitKarnam/KeyCloak/internal/utils/validators"
	storagebackendvalidator "github.com/AmitKarnam/KeyCloak/internal/utils/validators/storageBackendValidator"
	"github.com/AmitKarnam/KeyCloak/models"
)

type SecretEngineValidator struct{}

var _ validators.Validator[models.SecretEngine] = &SecretEngineValidator{}

// Validate if the secret engine model instance passed is valid or not.
func (sev *SecretEngineValidator) Validate(data models.SecretEngine) error {
	zapLogger.KeyCloaklogger.Infof("%v", "validating secret engine")

	// Need to check the validity of the Secret Engine Name
	if data.Name == "" {
		zapLogger.KeyCloaklogger.Errorf("%v", internalerrors.ErrEmptySecretEngineName)
		return errors.New("empty secret engine name value not allowed")
	}

	// Need to check the validity of the Secret Engine Encryption Strategy
	if data.Encryption_Strategy == "" {
		zapLogger.KeyCloaklogger.Errorf("%v", internalerrors.ErrInvalidEncryptionStrategy)
		return errors.New("empty secret engine encryption strategy value not allowed")
	}

	// Need to check the validity of the Secret Engine Storage Backend ( Call the Storage Backend Validator )
	storageValidator := storagebackendvalidator.StorageBackendValidator{}
	err := storageValidator.Validate(data.Storage_Backend)
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v", internalerrors.ErrInvalidStorageBackend)
		return err
	}

	zapLogger.KeyCloaklogger.Info("validation completed successfully")
	return nil
}
