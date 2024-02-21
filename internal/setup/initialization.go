package setup

import (
	"fmt"

	"github.com/jklaiber/jumper/internal/config"
	"github.com/jklaiber/jumper/pkg/inventory"
)

func Initialize(invReader inventory.InventoryReader, invParser inventory.InventoryParser) (*inventory.InventoryService, error) {
	if !config.ConfigurationFileExists() {
		fmt.Println("No configuration file found, setting up jumper...")
		if err := Setup(); err != nil {
			return nil, fmt.Errorf("could not setup jumper: %v", err)
		}
	}

	err := config.Parse()
	if err != nil {
		return nil, fmt.Errorf("could not initialize config: %v", err)
	}

	if !config.IsConfigured() {
		fmt.Println("Configuration is not complete, setting up jumper...")
		if err := Setup(); err != nil {
			return nil, fmt.Errorf("could not setup jumper: %v", err)
		}
	}

	invService, err := inventory.NewInventoryService(invReader, invParser)
	if err != nil {
		return nil, fmt.Errorf("could not create inventory service: %v", err)
	}

	return invService, nil
}
