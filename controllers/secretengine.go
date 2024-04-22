package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AmitKarnam/KeyCloak/database/sqlite"
	"github.com/AmitKarnam/KeyCloak/internal/internalerrors"
	"github.com/AmitKarnam/KeyCloak/internal/utils/logger/zapLogger"
	"github.com/AmitKarnam/KeyCloak/models"
)

type SecretEngineController struct{}

func (sec *SecretEngineController) Get(c *gin.Context) {
	// Connect to database
	conn, err := sqlite.DB.GetConnection()
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v , %v: %v", internalerrors.ErrConnectingToKCDB, "Unable to list secret engines", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": internalerrors.ErrConnectingToKCDB,
		})
		return
	}

	// Fetch all secret engines
	var scs []models.SecretEngine
	if err = conn.Find(&scs).Error; err != nil {
		// handle error
	}

	// Serialise them all into a JSON and return as response
}

func (sec *SecretEngineController) Post(c *gin.Context) {
	var se models.SecretEngine
	if err := c.BindJSON(&se); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(se.Name)
	fmt.Println(se.Encryption_Strategy)
	fmt.Printf("%s", se.Storage_Backend)
}
