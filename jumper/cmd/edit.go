package cmd

import (
	"log"

	"github.com/jklaiber/jumper/internal/config"
	"github.com/jklaiber/jumper/pkg/editor"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit files",
}

var editInvCmd = &cobra.Command{
	Use:   "inventory",
	Short: "Edit inventory file",
	PreRun: func(cmd *cobra.Command, args []string) {
		err := config.Parse()
		if err != nil {
			log.Fatalf("could not initialize config: %v", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := editor.EditInventory(config.Params.InventoryPath)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.AddCommand(editInvCmd)
}
