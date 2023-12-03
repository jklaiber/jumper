/*
Copyright © 2022 Julian Klaiber

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"log"

	"github.com/jklaiber/jumper/pkg/common"
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
	err := common.InitConfig()
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
