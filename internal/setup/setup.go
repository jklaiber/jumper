package setup

import (
	"fmt"
	"os"

	"github.com/jklaiber/jumper/internal/common"
	"github.com/jklaiber/jumper/internal/config"
	"github.com/jklaiber/jumper/internal/secret"
	"github.com/jklaiber/jumper/pkg/inventory"
	"github.com/manifoldco/promptui"
	vault "github.com/sosedoff/ansible-vault-go"
	"gopkg.in/yaml.v2"
)

const logo = `
    _                                 
   (_)                                
    _ _   _ _ __ ___  _ __   ___ _ __ 
   | | | | |  _   _ \|  _ \ / _ \  __|
   | | |_| | | | | | | |_) |  __/ |   
   | |\__ _|_| |_| |_|  __/ \___|_|   
  _/ |               | |              
 |__/                |_|              

`

type ConfigFile struct {
	InventoryFilePath string `yaml:"inventory_file"`
}

func confirm(text string) error {
	prompt := promptui.Prompt{
		Label:     text,
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		return fmt.Errorf("confirmation failed %v", err)
	}

	return nil
}

func promptSecret() (secret string, err error) {
	validate := func(input string) error {
		if len(input) < 6 {
			return fmt.Errorf("password must have more than 6 characters")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Encryption Password",
		Validate: validate,
		Mask:     '*',
	}

	secret, err = prompt.Run()
	if err != nil {
		return "", fmt.Errorf("password prompt failed %v", err)
	}

	return secret, nil
}

func promptInventoryDestination() (destination string, err error) {
	validate := func(input string) error {
		if len(input) < 1 {
			return fmt.Errorf("destination must be a valid path")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Inventory destination",
		Validate: validate,
	}

	destination, err = prompt.Run()

	if err != nil {
		return "", fmt.Errorf("inventory destination prompt failed %v", err)
	}
	return destination, nil
}

func createInventory(inventoryDestination string) error {
	defaultInventory := inventory.DefaultInventory()
	content, err := yaml.Marshal(&defaultInventory)
	if err != nil {
		return fmt.Errorf("could not marshal default inventory")
	}

	file, err := os.OpenFile(inventoryDestination, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("could not create inventory file")
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("could not close inventory file")
		}
	}()

	err = vault.EncryptFile(inventoryDestination, string(content), secret.GetSecretFromKeyring())
	if err != nil {
		return fmt.Errorf("could not encrypt inventory file")
	}

	return nil
}

func createConfigFile(inventoryDestination string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("could not get home directory")
	}

	configFile := ConfigFile{InventoryFilePath: inventoryDestination}
	data, err := yaml.Marshal(&configFile)
	if err != nil {
		return fmt.Errorf("could not marshal config file")
	}

	err = os.WriteFile(home+"/"+common.ConfigurationFileName, data, 0644)
	if err != nil {
		return fmt.Errorf("could not write config file")
	}

	return nil
}

func Setup() error {
	fmt.Print(logo)
	fmt.Println("It seems, that jumper is not fully configured.")
	fmt.Println("Please follow the steps to setup jumper:")
	fmt.Println("")
	if err := confirm("Do you want to configure jumper"); err != nil {
		return err
	}
	if !config.ConfigurationFileExists() {
		if err := confirm("Do you want to create a new configuration file"); err != nil {
			return err
		}
		inventoryDestination, err := promptInventoryDestination()
		if err != nil {
			return err
		}
		if err := createConfigFile(inventoryDestination); err != nil {
			return err
		}
	}
	if !secret.SecretAvailableFromKeyring() {
		if err := confirm("Do you want to create a new encryption password"); err != nil {
			return err
		}
		s, err := promptSecret()
		if err != nil {
			return err
		}
		secret.SetSecretInKeyring(s)
	}
	if !config.InventoryFileExists() {
		if err := config.Parse(); err != nil {
			return err
		}
		if err := confirm("Do you want to create a new empty inventory file"); err != nil {
			return err
		}

		inventory_path, err := config.GetInventoryFilePath()
		if err != nil {
			return err
		}

		if err := createInventory(inventory_path); err != nil {
			return err
		}
	}

	if err := confirm("Jumper is now configured. Do you want to continue"); err != nil {
		os.Exit(0)
	}

	return nil
}
