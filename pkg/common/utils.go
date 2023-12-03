package common

import (
	"fmt"
	"os"
	"strings"

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

func GetInventoryFilePath() (string, error) {
	inventory_path := viper.GetString(InventoryYamlKey)
	if strings.HasPrefix(inventory_path, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("could not get home directory")
		}
		inventory_path = home + inventory_path[1:]
	}
	return inventory_path, nil
}

func InventoryFileExists() bool {
	inventory_path, err := GetInventoryFilePath()
	if err != nil {
		return false
	}
	if _, err := os.Stat(inventory_path); os.IsNotExist(err) {
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
