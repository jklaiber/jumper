package cmd

import (
	"log"

	"github.com/jklaiber/jumper/internal/config"
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
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.Initialize(); err != nil {
			log.Fatalf("could not initialize jumper: %v", err)
		}

		accessConfig, err := config.Inv.GetAccessConfig(Group, args[0])
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
		return config.Inv.GetUngroupedHosts(), cobra.ShellCompDirectiveNoFileComp
	}
	return config.Inv.GetGroupHosts(Group), cobra.ShellCompDirectiveNoFileComp
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringVarP(&Group, groupFlagName, groupFlagNameShort, "", groupFlagDescription)
	if err := connectCmd.RegisterFlagCompletionFunc(groupFlagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return config.Inv.GetGroups(), cobra.ShellCompDirectiveDefault
	}); err != nil {
		log.Fatalf("could not register flag completion: %v", err)
	}
}
