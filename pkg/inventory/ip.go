package inventory

import "errors"

func (inventory *Inventory) GetIp(group string, host string) (string, error) {
	if group != "" {
		ip, err := inventory.getGroupHostIp(group, host)
		if err != nil {
			ip, err = inventory.getGroupIp(group)
			if err != nil {
				ip, err = inventory.getGlobalIp()
				if err != nil {
					return "", err
				}
				return ip, nil
			}
			return ip, nil
		}
		return ip, nil
	}
	ip, err := inventory.getUngroupedHostIp(host)
	if err != nil {
		ip, err = inventory.getGlobalIp()
		if err != nil {
			return "", err
		}
		return ip, nil
	}
	return ip, nil
}

func (inventory *Inventory) getGlobalIp() (string, error) {
	ip := inventory.All.Vars.Ip
	if ip == "" {
		return "", errors.New("global ip does not exist")
	}
	return ip, nil
}

func (inventory *Inventory) getGroupIp(group string) (string, error) {
	ip := inventory.All.Children[group].Vars.Ip
	if ip == "" {
		return "", errors.New("group ip does not exist")
	}
	return ip, nil
}

func (inventory *Inventory) getGroupHostIp(group string, host string) (string, error) {
	ip := inventory.All.Children[group].Hosts[host].Ip
	if ip == "" {
		return "", errors.New("host ip does not exist")
	}
	return ip, nil
}

func (inventory *Inventory) getUngroupedHostIp(host string) (string, error) {
	ip := inventory.All.Hosts[host].Ip
	if ip == "" {
		return "", errors.New("host ip does not exist")
	}
	return ip, nil
}
