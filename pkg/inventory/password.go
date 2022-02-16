package inventory

import "errors"

func (inventory *Inventory) GetPassword(group string, host string) (string, error) {
	if group != "" {
		password, err := inventory.getGroupHostPassword(group, host)
		if err != nil {
			password, err = inventory.getGroupPassword(group)
			if err != nil {
				password, err = inventory.getGlobalPassword()
				if err != nil {
					return "", err
				}
				return password, nil
			}
			return password, nil
		}
		return password, nil
	}
	password, err := inventory.getUngroupedHostPassword(host)
	if err != nil {
		password, err = inventory.getGlobalPassword()
		if err != nil {
			return "", err
		}
		return password, nil
	}
	return password, nil
}

func (inventory *Inventory) getGlobalPassword() (string, error) {
	password := inventory.All.Vars.Password
	if password == "" {
		return "", errors.New("global password does not exist")
	}
	return password, nil
}

func (inventory *Inventory) getGroupPassword(group string) (string, error) {
	password := inventory.All.Children[group].Vars.Password
	if password == "" {
		return "", errors.New("group password does not exist")
	}
	return password, nil
}

func (inventory *Inventory) getGroupHostPassword(group string, host string) (string, error) {
	password := inventory.All.Children[group].Hosts[host].Password
	if password == "" {
		return "", errors.New("host password does not exist")
	}
	return password, nil
}

func (inventory *Inventory) getUngroupedHostPassword(host string) (string, error) {
	password := inventory.All.Hosts[host].Password
	if password == "" {
		return "", errors.New("host password does not exist")
	}
	return password, nil
}
