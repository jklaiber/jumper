package inventory

import (
	"fmt"
)

func (inventory *Inventory) GetSshKey(group string, host string) (string, error) {
	if group != "" {
		if sshkey, err := inventory.getGroupHostSshKey(group, host); err == nil {
			return sshkey, nil
		}
		if sshkey, err := inventory.getGroupSshKey(group); err == nil {
			return sshkey, nil
		}
	}
	if sshkey, err := inventory.getUngroupedHostSshKey(host); err == nil {
		return sshkey, nil
	}
	return inventory.getGlobalSshKey()
}

func (inventory *Inventory) getGlobalSshKey() (string, error) {
	return getSshKeyFromVars(inventory.All.Vars)
}

func (inventory *Inventory) getGroupSshKey(group string) (string, error) {
	return getSshKeyFromVars(inventory.All.Children[group].Vars)
}

func (inventory *Inventory) getGroupHostSshKey(group string, host string) (string, error) {
	return getSshKeyFromVars(inventory.All.Children[group].Hosts[host])
}

func (inventory *Inventory) getUngroupedHostSshKey(host string) (string, error) {
	return getSshKeyFromVars(inventory.All.Hosts[host])
}

func getSshKeyFromVars(vars Vars) (string, error) {
	if vars.SshKey != "" {
		return vars.SshKey, nil
	}
	return "", fmt.Errorf("ssh key not enabled")
}
