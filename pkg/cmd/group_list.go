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
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

const (
	allFlagName        = "all"
	allFlagNameShort   = "A"
	allFlagDescription = "list all groups and its hosts"
)

var AllGroups bool

var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "list groups",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if AllGroups {
			printAllGroups()
		}
		if Name != "" {
			printSpecificGroup(Name)
		}
	},
}

func printSpecificGroup(group string) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{group})
	for _, host := range inv.GetGroupHosts(group) {
		t.AppendRow([]interface{}{host})
	}
	t.Render()
}

func printAllGroups() {
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"group", "hosts"}, rowConfigAutoMerge)
	for _, group := range inv.GetGroups() {
		for key, host := range inv.GetGroupHosts(group) {
			if key == 0 {
				t.AppendRow([]interface{}{group, host}, rowConfigAutoMerge)
			} else {
				t.AppendRow([]interface{}{"", host}, rowConfigAutoMerge)
			}
		}
		t.AppendSeparator()
	}
	t.Render()
}

func init() {
	groupCmd.AddCommand(groupListCmd)

	groupListCmd.Flags().BoolVarP(&AllGroups, allFlagName, allFlagNameShort, false, allFlagDescription)
	groupListCmd.Flags().StringVarP(&Name, NameFlagName, NameFlagNameShort, "", "list specific group and its hosts")
	groupListCmd.RegisterFlagCompletionFunc(NameFlagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return inv.GetGroups(), cobra.ShellCompDirectiveDefault
	})
}
