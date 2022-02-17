package inventory

import "errors"

func (inventory *Inventory) GetAddress(group string, host string) (string, error) {
	if group != "" {
		address, err := inventory.getGroupHostAddress(group, host)
		if err != nil {
			address, err = inventory.getGroupAddress(group)
			if err != nil {
				address, err = inventory.getGlobalAddress()
				if err != nil {
					return "", err
				}
				return address, nil
			}
			return address, nil
		}
		return address, nil
	}
	address, err := inventory.getUngroupedHostAddress(host)
	if err != nil {
		address, err = inventory.getGlobalAddress()
		if err != nil {
			return "", err
		}
		return address, nil
	}
	return address, nil
}

func (inventory *Inventory) getGlobalAddress() (string, error) {
	address := inventory.All.Vars.Address
	if address == "" {
		return "", errors.New("global address does not exist")
	}
	return address, nil
}

func (inventory *Inventory) getGroupAddress(group string) (string, error) {
	address := inventory.All.Children[group].Vars.Address
	if address == "" {
		return "", errors.New("group address does not exist")
	}
	return address, nil
}

func (inventory *Inventory) getGroupHostAddress(group string, host string) (string, error) {
	address := inventory.All.Children[group].Hosts[host].Address
	if address == "" {
		return "", errors.New("host address does not exist")
	}
	return address, nil
}

func (inventory *Inventory) getUngroupedHostAddress(host string) (string, error) {
	address := inventory.All.Hosts[host].Address
	if address == "" {
		return "", errors.New("host address does not exist")
	}
	return address, nil
}
