package inventory

import (
	"fmt"
)

func (inventory *Inventory) GetPassword(group string, host string) (string, error) {
	if group != "" {
		if password, err := inventory.getGroupHostPassword(group, host); err == nil {
			return password, nil
		}
		if password, err := inventory.getGroupPassword(group); err == nil {
			return password, nil
		}
	}
	if password, err := inventory.getUngroupedHostPassword(host); err == nil {
		return password, nil
	}
	return inventory.getGlobalPassword()
}

func (inventory *Inventory) getGlobalPassword() (string, error) {
	return getPasswordFromVars(inventory.All.Vars)
}

func (inventory *Inventory) getGroupPassword(group string) (string, error) {
	return getPasswordFromVars(inventory.All.Children[group].Vars)
}

func (inventory *Inventory) getGroupHostPassword(group string, host string) (string, error) {
	return getPasswordFromVars(inventory.All.Children[group].Hosts[host])
}

func (inventory *Inventory) getUngroupedHostPassword(host string) (string, error) {
	return getPasswordFromVars(inventory.All.Hosts[host])
}

func getPasswordFromVars(vars Vars) (string, error) {
	if vars.Password != "" {
		return vars.Password, nil
	}
	if vars.AnsibleSshPASS != "" {
		return vars.AnsibleSshPASS, nil
	}
	return "", fmt.Errorf("no password found for host")
}
