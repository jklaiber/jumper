package inventory

import "errors"

func (inventory *Inventory) GetUsername(group string, host string) (string, error) {
	if group != "" {
		username, err := inventory.getGroupHostUsername(group, host)
		if err != nil {
			username, err = inventory.getGroupUsername(group)
			if err != nil {
				username, err = inventory.getGlobalUsername()
				if err != nil {
					return "", err
				}
				return username, nil
			}
			return username, nil
		}
		return username, nil
	}
	username, err := inventory.getUngroupedHostUsername(host)
	if err != nil {
		username, err = inventory.getGlobalUsername()
		if err != nil {
			return "", err
		}
		return username, nil
	}
	return username, nil
}

func (inventory *Inventory) getGlobalUsername() (string, error) {
	username := inventory.All.Vars.Username
	if username == "" {
		return "", errors.New("global username does not exist")
	}
	return username, nil
}

func (inventory *Inventory) getGroupUsername(group string) (string, error) {
	username := inventory.All.Children[group].Vars.Username
	if username == "" {
		return "", errors.New("group username does not exist")
	}
	return username, nil
}

func (inventory *Inventory) getGroupHostUsername(group string, host string) (string, error) {
	username := inventory.All.Children[group].Hosts[host].Username
	if username == "" {
		return "", errors.New("host username does not exist")
	}
	return username, nil
}

func (inventory *Inventory) getUngroupedHostUsername(host string) (string, error) {
	username := inventory.All.Hosts[host].Username
	if username == "" {
		return "", errors.New("host username does not exist")
	}
	return username, nil
}
