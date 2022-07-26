package inventory

import "errors"

func (inventory *Inventory) GetAddress(group string, host string) (string, error) {
	if group != "" {
		address, err := inventory.getGroupHostAddress(group, host)
		if err != nil {
			return address, nil
		}
		return address, nil
	}
	address, err := inventory.getUngroupedHostAddress(host)
	if err != nil {
		return address, nil
	}
	return address, nil
}

func (inventory *Inventory) getGroupHostAddress(group string, host string) (string, error) {
	address := ""
	if inventory.All.Children[group].Hosts[host].Address != "" {
		address = inventory.All.Children[group].Hosts[host].Address
	} else {
		address = inventory.All.Children[group].Hosts[host].AnsibleHost
	}
	if address == "" {
		return "", errors.New("host address does not exist")
	}
	return address, nil
}

func (inventory *Inventory) getUngroupedHostAddress(host string) (string, error) {
	address := ""
	if inventory.All.Hosts[host].Address != "" {
		address = inventory.All.Hosts[host].Address
	} else {
		address = inventory.All.Hosts[host].AnsibleHost
	}
	if address == "" {
		return "", errors.New("host address does not exist")
	}
	return address, nil
}
