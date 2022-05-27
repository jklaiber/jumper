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
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var groupAddCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new group",
	// Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// createNewNote()
	},
}

func init() {
	groupCmd.AddCommand(groupAddCmd)

	// groupAddCmd.Flags().StringVarP(&Name, NameFlagName, NameFlagNameShort, "", "name of the new group")
	// groupAddCmd.MarkFlagRequired(NameFlagName)

	// groupAddCmd.Flags().StringVarP(&Password, PasswordFlagName, PasswordFlagShort, "", PasswordFlagDescription)

	// groupAddCmd.Flags().StringVarP(&Username, UsernameFlagName, UsernameFlagShort, "", UsernameFlagDescription)

	// groupAddCmd.Flags().StringVarP(&Key, KeyFlagName, KeyFlagShort, "", KeyFlagDescription)

	// groupAddCmd.Flags().BoolVarP(&Agent, AgentFlagName, AgentFlagShort, false, AgentFlagDescription)
}

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func createNewNote() {
	hostNamePromptContent := promptContent{
		"Please provide a host name.",
		"What is the name of the new host?",
	}
	hostName := promptGetInput(hostNamePromptContent)

	// belongsToGroupPromptContent := promptContent{
	// 	"Please provide if the host belongs to a group.",
	// 	fmt.Sprintf("Belongs the host: %s to a group?", hostName),
	// }
	// belongsToGroup := promptGetInput(belongsToGroupPromptContent)

	belongsToGroupPrompt := promptui.Prompt{
		Label:     "Belongs the host: to a group?",
		IsConfirm: false,
	}

	belongsToGroup, err := belongsToGroupPrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Println(belongsToGroup)

	categoryPromptContent := promptContent{
		"Please provide the group to which the host belongs.",
		fmt.Sprintf("To which group does the host: %s belongs?", hostName),
	}
	category := promptGetSelect(categoryPromptContent)

	// data.insertnote(word, definition, category)

	fmt.Println(hostName)
	// fmt.Println(definition)
	fmt.Println(category)
}
