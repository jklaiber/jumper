package cmd

import (
	"fmt"

	"github.com/jklaiber/jumper/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of jumper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Jumper version", version.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
