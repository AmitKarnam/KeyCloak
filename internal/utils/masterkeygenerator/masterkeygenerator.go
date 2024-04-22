package masterkeygenerator

import (
	"crypto/rand"
	"os"

	internalCrypto "github.com/AmitKarnam/KeyCloak/internal/utils/crypto"
	"github.com/AmitKarnam/KeyCloak/internal/utils/logger/zapLogger"
)

func MasterKeyHandler() error {
	masterKey, err := generateMasterKey()
	if err != nil {
		return err
	}

	err = storeMasterKey(masterKey)
	if err != nil {
		return err
	}

	return nil
}

func generateMasterKey() ([]byte, error) {
	// Generate's a new 32 byte master key required by the choosen encryption algorithm ( ChaCha20-Poly1305 ).
	key := make([]byte, internalCrypto.KeySize)
	_, err := rand.Read(key)
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("Error generating master key : %v", err)
		return nil, err
	}

	zapLogger.KeyCloaklogger.Infof("Generated master key successfully")
	return key, nil
}

func storeMasterKey(masterKey []byte) error {
	// Store the generated master key in the file location. Opens an existing file, if not present throw error

	// path should be fetched from config file, hardcoded as of now
	f, err := os.OpenFile("./internal/random.txt", os.O_RDWR, 0666)
	if err != nil {
		// thinking of not passing the error into the logs as it gives away the destination of the file.
		zapLogger.KeyCloaklogger.Errorf("Error opening file to store master key : %v", err)
		return err
	}
	defer f.Close()

	_, err = f.Write(masterKey)
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("Error writing master key into the file : %v", err)
		return err
	}

	zapLogger.KeyCloaklogger.Infof("Successfully stored master key into the file.")
	return nil
}
