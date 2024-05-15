package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AmitKarnam/KeyCloak/database/sqlite"
	"github.com/AmitKarnam/KeyCloak/internal/internalerrors"
	"github.com/AmitKarnam/KeyCloak/internal/utils/logger/zapLogger"
	SEValidator "github.com/AmitKarnam/KeyCloak/internal/utils/validators/secretEngineValidator"
	"github.com/AmitKarnam/KeyCloak/models"
)

type SecretEngineController struct{}

// Get controller is invoked when a GET request on the secretengine is triggered. It lists out all the secret engine configured.
func (sec *SecretEngineController) Get(c *gin.Context) {
	// Connect to database
	conn, err := sqlite.DB.GetConnection()
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v , %v: %v", internalerrors.ErrConnectingToKCDB, "Unable to list secret engines", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": internalerrors.ErrConnectingToKCDB.Error(),
		})
		return
	}

	// Fetching all secret engines
	var scs []models.SecretEngine
	if err = conn.Find(&scs).Error; err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v: %v", "Unable to list secret engines", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, scs)
}

// Post controller is invoked when a POST request is received for the secretengine endpoint. It is used to register a new secret engine.
func (sec *SecretEngineController) Post(c *gin.Context) {

	// In-Progress: Need to handle case of partial data in request body.
	// Once I bind it to the variable, I'll have to check for all attributes of the secret engine model.
	var se models.SecretEngine
	if err := c.BindJSON(&se); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call Validate on the se variable to validate the secret engine type
	var secretEngnineValidation SEValidator.SecretEngineValidator
	err := secretEngnineValidation.Validate(se)
	if err != nil {
		//log error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	conn, err := sqlite.DB.GetConnection()
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v , %v: %v", internalerrors.ErrConnectingToKCDB, "Unable to post secret engine record", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": internalerrors.ErrConnectingToKCDB.Error(),
		})
		return
	}

	// Check for duplicate entries, No duplicate entries wrt to secret name is allowed
	var existingRecord models.SecretEngine
	if err = conn.Where("name = ?", se.Name).First(&existingRecord).Error; err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if existingRecord.ID != 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "Secret engine with the same name already exists"})
		return
	}

	if err = conn.Create(&se).Error; err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v: %v", "Unable to post secret engine record", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "New secret engine created successfuly"})
}

// Delete controller is invoked when a DELETE request is recieved to the secretengine endpoint. It is used to delete the required secret engine
func (sec *SecretEngineController) Delete(c *gin.Context) {
	scDelete := c.Param("name")

	conn, err := sqlite.DB.GetConnection()
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v , %v: %v", internalerrors.ErrConnectingToKCDB, "Unable to delete secret engine record", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": internalerrors.ErrConnectingToKCDB.Error(),
		})
		return
	}

	var secretEngine models.SecretEngine
	if err = conn.Where("name=?", scDelete).Find(&secretEngine).Error; err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if secretEngine.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Secret engine not found"})
		return
	}

	if err = conn.Delete(&secretEngine).Error; err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v: %v", "Unable to delete secret engine record", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": internalerrors.ErrConnectingToKCDB.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Secret engine deleted successfully"})

}
