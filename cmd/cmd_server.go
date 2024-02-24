/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"github.com/AmitKarnam/KeyCloak/database/sqlite"
	"github.com/AmitKarnam/KeyCloak/internal/internalerrors"
	"github.com/AmitKarnam/KeyCloak/internal/utlis/logger/zapLogger"
	models "github.com/AmitKarnam/KeyCloak/models"
	"github.com/AmitKarnam/KeyCloak/server"
)

// cmdServerCmd represents the cmdServer command
var cmdServerCmd = &cobra.Command{
	Use:   "server",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {

		//
		var wg = new(errgroup.Group)

		// Initialize Logger
		zlog := zapLogger.ZapLogging{}
		zlog.GenerateLogger()

		// Initialize DB
		// Migrate DB
		err := initializeDB()
		if err != nil {
			return err
		}

		// Initialize REST API
		wg.Go(func() error {
			return runServer()
		})
		zapLogger.KeyCloaklogger.Infof("Successfully running KeyCloak REST Server.")

		// To Do
		// Initialize the master key scheduler. As a part of initialisation it should generate a
		// master key and store it in the desired location.

		if err := wg.Wait(); err != nil {
			return err
		}

		return nil

	},
}

// Initialize DB is an asyncronous call, it inherently calls migrateDB
func initializeDB() error {
	// Hardcoded , should be derived from configuration file
	sqlite.DBInit("keycloak.db")

	err := migrateDB()
	if err != nil {
		return err
	}
	zapLogger.KeyCloaklogger.Infof("Successfully Initialized and Migrated Key Cloak DB.")

	return nil
}

func migrateDB() error {
	db, err := sqlite.DB.GetConnection()
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v while migrating: %v", internalerrors.ErrConnectingToKCDB, err)
		return err
	}

	err = db.AutoMigrate(models.SecretEngine{})
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v : %v", internalerrors.ErrMigratingKCDB, err)
		return err
	}

	return nil
}

func runServer() error {
	err := server.RunServer()
	if err != nil {
		zapLogger.KeyCloaklogger.Errorf("%v : %v", internalerrors.ErrRunningServer, err)
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(cmdServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmdServerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmdServerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
