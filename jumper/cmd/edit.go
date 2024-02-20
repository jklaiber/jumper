package cmd

import (
	"log"

	"github.com/jklaiber/jumper/internal/config"
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
		if err := config.Initialize(); err != nil {
			log.Fatalf("could not initialize jumper: %v", err)
		}
		err := config.Inv.EditInventory(viper.GetString("inventory_file"))
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.AddCommand(editInvCmd)
}
