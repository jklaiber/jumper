package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/jklaiber/jumper/internal/config"
)

func ConfigurationFileExists() bool {
	configurationFilePath, err := config.GetConfigurationFilePath()
	if err != nil {
		return false
	}
	if _, err := os.Stat(configurationFilePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetInventoryFilePath() (string, error) {
	inventory_path := config.Params.InventoryPath
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
