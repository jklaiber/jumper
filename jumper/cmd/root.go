package cmd

import (
	"log"

	"github.com/jklaiber/jumper/internal/common"
	"github.com/jklaiber/jumper/internal/config"
	"github.com/jklaiber/jumper/pkg/inventory"
	"github.com/jklaiber/jumper/pkg/setup"
	"github.com/spf13/cobra"
)

var inv inventory.Inventory

var rootCmd = &cobra.Command{
	Use:   "jumper",
	Short: "A simple cli SSH manager",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	err := config.Parse()
	if err != nil {
		log.Fatalf("could not initialize config: %v", err)
	}

	if !common.IsConfigured() {
		if err := setup.Setup(); err != nil {
			log.Fatalf("could not setup jumper: %v", err)
		}
	}
	cobra.OnInitialize(initInventory)
}

func initInventory() {
	inventoryFile, err := common.GetInventoryFilePath()
	if err != nil {
		log.Fatalf("could not get inventory file path")
	}
	inventory, err := inventory.NewInventory(inventoryFile)
	if err != nil {
		log.Fatalf("could not create inventory")
	}
	inv = inventory
}
