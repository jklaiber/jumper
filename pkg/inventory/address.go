package inventory

import "errors"

func (inventory *Inventory) GetAddress(group string, host string) (string, int, error) {
	if group != "" {
		address, port, err := inventory.getGroupHostAddress(group, host)
		if err != nil {
			return address, port, nil
		}
		return address, port, nil
	}
	address, port, err := inventory.getUngroupedHostAddress(host)
	if err != nil {
		return address, port, nil
	}
	return address, port, nil
}

func (inventory *Inventory) getGroupHostAddress(group string, host string) (string, int, error) {
	address := ""
	if inventory.All.Children[group].Hosts[host].Address != "" {
		address = inventory.All.Children[group].Hosts[host].Address
	} else {
		address = inventory.All.Children[group].Hosts[host].AnsibleHost
	}
	if address == "" {
		return "", 0, errors.New("host address does not exist")
	}
	port := inventory.All.Children[group].Hosts[host].Port
	return address, port, nil
}

func (inventory *Inventory) getUngroupedHostAddress(host string) (string, int, error) {
	address := ""
	if inventory.All.Hosts[host].Address != "" {
		address = inventory.All.Hosts[host].Address
	} else {
		address = inventory.All.Hosts[host].AnsibleHost
	}
	if address == "" {
		return "", 0, errors.New("host address does not exist")
	}
	port := inventory.All.Hosts[host].Port
	return address, port, nil
}
