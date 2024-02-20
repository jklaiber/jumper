package inventory

import (
	"fmt"
)

func (inventory *Inventory) GetAddress(group string, host string) (string, int, error) {
	if group != "" {
		return inventory.getGroupHostAddress(group, host)
	}
	return inventory.getUngroupedHostAddress(host)
}

func (inventory *Inventory) getGroupHostAddress(group string, host string) (string, int, error) {
	return inventory.getHostAddress(inventory.All.Children[group].Hosts[host])
}

func (inventory *Inventory) getUngroupedHostAddress(host string) (string, int, error) {
	return inventory.getHostAddress(inventory.All.Hosts[host])
}

func (inventory *Inventory) getHostAddress(vars Vars) (string, int, error) {
	address := vars.Address
	if address == "" {
		address = vars.AnsibleHost
	}
	if address == "" {
		return "", 0, fmt.Errorf("no address found for host")
	}
	return address, vars.Port, nil
}
