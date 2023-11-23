package inventory

import (
	"fmt"
)

func (inventory *Inventory) GetUsername(group string, host string) (string, error) {
	if group != "" {
		if username, err := inventory.getGroupHostUsername(group, host); err == nil {
			return username, nil
		}
		if username, err := inventory.getGroupUsername(group); err == nil {
			return username, nil
		}
	}
	if username, err := inventory.getUngroupedHostUsername(host); err == nil {
		return username, nil
	}
	return inventory.getGlobalUsername()
}

func (inventory *Inventory) getGlobalUsername() (string, error) {
	return getUsernameFromVars(inventory.All.Vars)
}

func (inventory *Inventory) getGroupUsername(group string) (string, error) {
	return getUsernameFromVars(inventory.All.Children[group].Vars)
}

func (inventory *Inventory) getGroupHostUsername(group string, host string) (string, error) {
	return getUsernameFromVars(inventory.All.Children[group].Hosts[host])
}

func (inventory *Inventory) getUngroupedHostUsername(host string) (string, error) {
	return getUsernameFromVars(inventory.All.Hosts[host])
}

func getUsernameFromVars(vars Vars) (string, error) {
	if vars.Username != "" {
		return vars.Username, nil
	}
	if vars.AnsibleUser != "" {
		return vars.AnsibleUser, nil
	}
	return "", fmt.Errorf("no username found for host")
}
