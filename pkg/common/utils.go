package common

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func InitConfig() error {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".jumper")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("could not read config file: %v", err)
	}

	return nil
}

func GetConfigurationFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home directory")
	}
	return home + "/" + ConfigurationFileName, nil
}

func ConfigurationFileExists() bool {
	configurationFilePath, err := GetConfigurationFilePath()
	if err != nil {
		return false
	}
	if _, err := os.Stat(configurationFilePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetInventoryFilePath() string {
	return viper.GetString(InventoryYamlKey)
}

func InventoryFileExists() bool {
	if _, err := os.Stat(GetInventoryFilePath()); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsConfigured() bool {
	if ConfigurationFileExists() && InventoryFileExists() && SecretAvailableFromKeyring() {
		return true
	}
	return false
}
