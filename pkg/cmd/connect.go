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

	"github.com/jklaiber/jumper/pkg/connection"
	"github.com/spf13/cobra"
)

const (
	groupFlagName        = "group"
	groupFlagNameShort   = "g"
	groupFlagDescription = "Connect to a host in a group"
)

var Group string

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a saved connection",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		accessConfig, err := inv.GetAccessConfig(Group, args[0])
		if err != nil {
			log.Fatal(err)
		}

		connection, err := connection.NewConnection(accessConfig)
		if err != nil {
			log.Fatal(err)
		}
		connection.Start()
	},
	ValidArgsFunction: UngroupedHostGet,
}

func UngroupedHostGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if Group == "" {
		return inv.GetUngroupedHosts(), cobra.ShellCompDirectiveNoFileComp
	}
	return inv.GetGroupHosts(Group), cobra.ShellCompDirectiveNoFileComp
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringVarP(&Group, groupFlagName, groupFlagNameShort, "", groupFlagDescription)
	if err := connectCmd.RegisterFlagCompletionFunc(groupFlagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return inv.GetGroups(), cobra.ShellCompDirectiveDefault
	}); err != nil {
		log.Fatalf("could not register flag completion: %v", err)
	}

}
