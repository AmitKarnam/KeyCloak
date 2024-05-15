package internalerrors

import (
	"errors"
)

var ErrConnectingToKCDB = errors.New("error connecting to KeyCloak database")
var ErrMigratingKCDB = errors.New("error migrating KeyCloak database")
var ErrRunningServer = errors.New("error encountered while running REST server")
var ErrEmptySecretEngineName = errors.New("empty secret engine name not allowed")
var ErrInvalidEncryptionStrategy = errors.New("invalid encryption strategy")
var ErrInvalidStorageBackend = errors.New("invalid storage backend")
