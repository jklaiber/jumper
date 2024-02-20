package config

import (
	"fmt"
	"os"

	"github.com/jklaiber/jumper/internal/common"
	"github.com/spf13/viper"
)

var (
	viperInstance = viper.New()
	Params        Config
)

type Config struct {
	InventoryPath string `mapstructure:"inventory_file"`
}

func init() {
	viperInstance.SetConfigName(".jumper")
	viperInstance.SetConfigType("yaml")
	viperInstance.AddConfigPath("$HOME")
}

func GetConfigurationFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home directory")
	}
	return home + "/" + common.ConfigurationFileName, nil
}

func Parse() error {
	if err := viperInstance.ReadInConfig(); err != nil {
		return fmt.Errorf("could not read config file: %v", err)
	}

	if err := viperInstance.Unmarshal(&Params); err != nil {
		return fmt.Errorf("could not unmarshal config: %v", err)
	}

	return nil
}

func GetInstance() *viper.Viper {
	return viperInstance
}
