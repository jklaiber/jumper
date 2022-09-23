package common

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func InitConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".jumper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("could not read config file")
	}
}

func GetConfigurationFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("could not get home directory")
	}
	return home + "/" + ConfigurationFileName
}

func ConfigurationFileExists() bool {
	if _, err := os.Stat(GetConfigurationFilePath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetInventoryFilePath() string {
	InitConfig()
	return viper.GetString(InventoryYamlKey)
}

func InventoryFileExists() bool {
	if _, err := os.Stat(GetInventoryFilePath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsConfigured() bool {
	if ConfigurationFileExists() {
		if InventoryFileExists() {
			if SecretAvailableFromKeyring() {
				return true
			}
		}
	}
	return false
}
