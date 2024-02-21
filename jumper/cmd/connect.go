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
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		accessConfig, err := invService.GetAccessConfig(Group, args[0])
		if err != nil {
			log.Fatalf("could not get access config: %v", err)
		}
		connection, err := connection.NewConnection(accessConfig)
		if err != nil {
			log.Fatalf("could not create connection: %v", err)
		}
		connection.Start()
	},
	ValidArgsFunction: UngroupedHostGet,
}

func UngroupedHostGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if Group == "" {
		hostDetails := invService.GetUngroupedHosts()
		hosts := make([]string, len(hostDetails))
		for i, host := range hostDetails {
			hosts[i] = host.Name
		}
		return hosts, cobra.ShellCompDirectiveNoFileComp
	}
	hostDetails := invService.GetGroupHosts(Group)
	hosts := make([]string, len(hostDetails))
	for i, host := range hostDetails {
		hosts[i] = host.Name
	}
	return hosts, cobra.ShellCompDirectiveNoFileComp
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringVarP(&Group, groupFlagName, groupFlagNameShort, "", groupFlagDescription)
	if err := connectCmd.RegisterFlagCompletionFunc(groupFlagName, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		groups := getGroupsAsSlice()
		return groups, cobra.ShellCompDirectiveDefault
	}); err != nil {
		log.Fatalf("could not register flag completion: %v", err)
	}
}

func getGroupsAsSlice() []string {
	groupDetails := invService.GetGroups()
	groups := make([]string, len(groupDetails))
	for i, group := range groupDetails {
		groups[i] = group.Name
	}
	return groups
}
