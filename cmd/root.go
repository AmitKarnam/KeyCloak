/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "keycloak",
	Short: "Manage your encrypted secrets with Keycloak",
	Long: `Keycloak is a powerful tool for managing encrypted secrets such as passwords, API keys, cloud credentials, and more. 

With Keycloak, you can securely store your secrets in the storage backend of your choice, thanks to its versatile secret engine. This engine not only encrypts your secrets but also handles their storage, providing an extra layer of security for your sensitive information.

The CLI tool provides an easy and efficient way to interact with Keycloak.

Keycloak is designed to make secret management as straightforward and secure as possible. Start using Keycloak today and take the first step towards better secret management.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.KeyCloak.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
