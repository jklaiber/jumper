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
	"os"

	"github.com/jklaiber/grasshopper/pkg/inventory"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var inv inventory.Inventory

var rootCmd = &cobra.Command{
	Use:   "grasshopper",
	Short: "A simple cli SSH manager",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initInventory)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.grasshopper.yaml)")

}

func initInventory() {
	inventoryFile := viper.GetString("inventory_file")
	inventoryPassword := viper.GetString("vault_password")
	inventory, err := inventory.NewInventory(inventoryFile, inventoryPassword)
	if err != nil {
		log.Fatalf("could not create inventory")
	}
	inv = inventory
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".grasshopper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("could not read config file")
	}
}
