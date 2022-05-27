package cmd

import (
	"github.com/spf13/cobra"
)

const (
	GroupFlagName             = "group"
	GroupFlagNameShort        = "g"
	GroupFlagDescription      = "connect to a host in a group"
	NameFlagName              = "name"
	NameFlagNameShort         = "n"
	ConnectionFlagName        = "connection"
	ConnectionFlagShort       = "c"
	ConnectionFlagDescription = "host connection address"
	PasswordFlagName          = "password"
	PasswordFlagShort         = "p"
	PasswordFlagDescription   = "login-password"
	UsernameFlagName          = "username"
	UsernameFlagShort         = "u"
	UsernameFlagDescription   = "username for login"
	KeyFlagName               = "key"
	KeyFlagShort              = "k"
	KeyFlagDescription        = "ssh-key to be used"
	AgentFlagName             = "ssh-agent"
	AgentFlagShort            = "a"
	AgentFlagDescription      = "use an ssh-agent"
)

var (
	Group      string
	Name       string
	Connection string
	Password   string
	Username   string
	Key        string
	Agent      bool
)

func UngroupedHostGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if Group == "" {
		return inv.GetUngroupedHosts(), cobra.ShellCompDirectiveNoFileComp
	}
	return inv.GetGroupHosts(Group), cobra.ShellCompDirectiveNoFileComp
}
