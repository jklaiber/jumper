package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit files",
}

var editInvCmd = &cobra.Command{
	Use:   "inventory",
	Short: "Edit inventory file",
	Run: func(cmd *cobra.Command, args []string) {
		err := inv.EditInventory(viper.GetString("inventory_file"))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.AddCommand(editInvCmd)
}
