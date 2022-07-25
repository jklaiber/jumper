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
	password := ""
	if inventory.All.Vars.Password != "" {
		password = inventory.All.Vars.Password
	} else {
		password = inventory.All.Vars.AnsibleSshPASS
	}
	if password == "" {
		return "", errors.New("global password does not exist")
	}
	return password, nil
}

func (inventory *Inventory) getGroupPassword(group string) (string, error) {
	password := ""
	if inventory.All.Children[group].Vars.Password != "" {
		password = inventory.All.Children[group].Vars.Password
	} else {
		password = inventory.All.Children[group].Vars.AnsibleSshPASS
	}
	if password == "" {
		return "", errors.New("group password does not exist")
	}
	return password, nil
}

func (inventory *Inventory) getGroupHostPassword(group string, host string) (string, error) {
	password := ""
	if inventory.All.Children[group].Hosts[host].Password != "" {
		password = inventory.All.Children[group].Hosts[host].Password
	} else {
		password = inventory.All.Children[group].Hosts[host].AnsibleSshPASS
	}
	if password == "" {
		return "", errors.New("host password does not exist")
	}
	return password, nil
}

func (inventory *Inventory) getUngroupedHostPassword(host string) (string, error) {
	password := ""
	if inventory.All.Hosts[host].Password != "" {
		password = inventory.All.Hosts[host].Password
	} else {
		password = inventory.All.Hosts[host].AnsibleSshPASS
	}
	if password == "" {
		return "", errors.New("host password does not exist")
	}
	return password, nil
}
