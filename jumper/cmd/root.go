package cmd

import (
	"log"
	"os"

	"github.com/jklaiber/jumper/internal/setup"
	"github.com/jklaiber/jumper/pkg/inventory"
	"github.com/spf13/cobra"
)

var invService inventory.InventoryManager

var rootCmd = &cobra.Command{
	Use:   "jumper",
	Short: "A simple cli SSH manager",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	if os.Getenv("JUMPER_SKIP_INV_INIT") == "" {
		invReader := inventory.DefaultInventoryReader{}
		invParser := inventory.DefaultInventoryParser{}
		inv, err := setup.Initialize(&invReader, &invParser)
		if err != nil {
			log.Fatalf("could not initialize jumper: %v", err)
		}
		invService = inv
	}
}
