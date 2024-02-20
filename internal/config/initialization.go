package config

import (
	"fmt"
	"log"

	"github.com/jklaiber/jumper/pkg/inventory"
)

var Inv inventory.Inventory

func Initialize() error {
	if !ConfigurationFileExists() {
		if err := Setup(); err != nil {
			return fmt.Errorf("could not setup jumper: %v", err)
		}
	}

	err := Parse()
	if err != nil {
		return fmt.Errorf("could not initialize config: %v", err)
	}

	if !IsConfigured() {
		if err := Setup(); err != nil {
			log.Fatalf("could not setup jumper: %v", err)
		}
	}

	err = initializeInventory()
	if err != nil {
		return err
	}

	return nil
}

func initializeInventory() error {
	inventoryFile, err := GetInventoryFilePath()
	if err != nil {
		return fmt.Errorf("could not get inventory file path")
	}
	inventory, err := inventory.NewInventory(inventoryFile)
	if err != nil {
		return fmt.Errorf("could not create inventory")
	}

	Inv = inventory

	return nil
}
