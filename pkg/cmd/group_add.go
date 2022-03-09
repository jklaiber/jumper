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
	"github.com/spf13/cobra"
)

var groupAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new group",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	groupCmd.AddCommand(groupAddCmd)

	groupAddCmd.Flags().StringVarP(&Name, NameFlagName, NameFlagNameShort, "", "name of the new group")
	groupAddCmd.MarkFlagRequired(NameFlagName)

	groupAddCmd.Flags().StringVarP(&Password, PasswordFlagName, PasswordFlagShort, "", PasswordFlagDescription)
	groupAddCmd.MarkFlagRequired(PasswordFlagName)

	groupAddCmd.Flags().StringVarP(&Username, UsernameFlagName, UsernameFlagShort, "", UsernameFlagDescription)
	groupAddCmd.MarkFlagRequired(UsernameFlagName)

	groupAddCmd.Flags().StringVarP(&Connection, ConnectionFlagName, ConnectionFlagShort, "", ConnectionFlagDescription)
	groupAddCmd.MarkFlagRequired(ConnectionFlagName)

	groupAddCmd.Flags().StringVarP(&Key, KeyFlagName, KeyFlagShort, "", KeyFlagDescription)
	groupAddCmd.MarkFlagRequired(KeyFlagName)

	groupAddCmd.Flags().BoolVarP(&Agent, AgentFlagName, AgentFlagShort, false, AgentFlagDescription)
	groupAddCmd.MarkFlagRequired(AgentFlagName)
}
