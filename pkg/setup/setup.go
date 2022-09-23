package setup

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jklaiber/jumper/pkg/common"
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

func Confirm(text string) {
	prompt := promptui.Prompt{
		Label:     text,
		IsConfirm: true,
	}

	_, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt aborted")
	}
}

func promptSecret() (secret string, err error) {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("Password must have more than 6 characters")
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
		log.Fatalf("Prompt failed %v\n", err)
		return
	}

	return secret, nil
}

func promptInventoryDestination() (destination string, err error) {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("Destination must be a valid path")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "Inventory destination",
		Validate: validate,
	}

	destination, err = prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return destination, nil
}

func createInventory(inventoryDestination string) {
	file, err := os.OpenFile(inventoryDestination, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		log.Fatalf("could not create empty inventory file")
	}
	file.Close()
	err = vault.EncryptFile(inventoryDestination, "", common.GetSecretFromKeyring())
	if err != nil {
		log.Fatalf("could not encrypt empty inventory file")
	}
}

func createConfigFile(inventoryDestination string) {
	home, err := os.UserHomeDir()
	config := ConfigFile{InventoryFilePath: inventoryDestination}
	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("could not marshal config file")
	}
	err = ioutil.WriteFile(home+"/"+common.ConfigurationFileName, data, 0644)
	if err != nil {
		log.Fatalf("could not write config file")
	}
}

func Setup() (err error) {
	fmt.Printf(logo)
	fmt.Println("It seems, that jumper is not fully configured")
	fmt.Println("Please follow the steps to setup jumper")
	fmt.Println("")
	Confirm("Do you want to configure jumper")
	if !common.ConfigurationFileExists() {
		Confirm("Do you want to create a new configuration file")
		inventoryDestination, err := promptInventoryDestination()
		if err != nil {
			return err
		}
		createConfigFile(inventoryDestination)
	}
	if !common.SecretAvailableFromKeyring() {
		Confirm("Do you want to create a new encryption password")
		secret, err := promptSecret()
		if err != nil {
			return err
		}
		common.SetSecretInKeyring(secret)
	}
	if !common.InventoryFileExists() {
		common.InitConfig()
		Confirm("Do you want to create a new empty inventory file")
		createInventory(common.GetInventoryFilePath())
	}
	os.Exit(0)
	return
}
