/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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

	"github.com/jklaiber/grasshopper/internal/connection"
	"github.com/jklaiber/grasshopper/pkg/inventory"
	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a saved connection",
	Long:  `A longer description that spans multiple lines an.`,
	Run: func(cmd *cobra.Command, args []string) {

		// filePath := "inventory.yaml"
		// b, err := ioutil.ReadFile(filePath)
		// if err != nil {
		// 	log.Fatalf("error")
		// }
		// mappedData := make(map[interface{}]interface{})
		// err = yaml.Unmarshal([]byte(b), &mappedData)
		// if err != nil {
		// 	log.Fatalf("error: %v", err)
		// }

		// fmt.Println(mappedData)

		// fmt.Println(mappedData["all"])

		// filePath = expandFilePath(filePath)
		inventory, err := inventory.NewInventory("inventory.yaml")
		if err != nil {
			log.Fatalf("could not create inventory")
		}
		// username, _ := inventory.GetUsername("automation", "awx")
		// password, _ := inventory.GetPassword("automation", "awx")
		// host := "awx"

		u, p, _, i, err := inventory.GetAccessInformation("automation", "awx")

		err = connection.NewConnection(u, i, p)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// connectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// connectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
