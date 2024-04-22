package internalerrors

import (
	"errors"
)

var ErrConnectingToKCDB = errors.New("Error connecting to Key Cloak database.")
var ErrMigratingKCDB = errors.New("Error migrating KeyCloak database.")
var ErrRunningServer = errors.New("Error encountered while running REST server.")
