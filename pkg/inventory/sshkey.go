package inventory

import "errors"

func (inventory *Inventory) GetSshKey(group string, host string) (string, error) {
	if group != "" {
		sshkey, err := inventory.getGroupHostSshKey(group, host)
		if err != nil {
			sshkey, err = inventory.getGroupSshKey(group)
			if err != nil {
				sshkey, err = inventory.getGlobalSshKey()
				if err != nil {
					return "", err
				}
				return sshkey, nil
			}
			return sshkey, nil
		}
		return sshkey, nil
	}
	sshkey, err := inventory.getUngroupedHostSshKey(host)
	if err != nil {
		sshkey, err = inventory.getGlobalSshKey()
		if err != nil {
			return "", err
		}
		return sshkey, nil
	}
	return sshkey, nil
}

func (inventory *Inventory) getGlobalSshKey() (string, error) {
	sshkey := inventory.All.Vars.SshKey
	if sshkey == "" {
		return "", errors.New("global sshkey does not exist")
	}
	return sshkey, nil
}

func (inventory *Inventory) getGroupSshKey(group string) (string, error) {
	sshkey := inventory.All.Children[group].Vars.SshKey
	if sshkey == "" {
		return "", errors.New("group sshkey does not exist")
	}
	return sshkey, nil
}

func (inventory *Inventory) getGroupHostSshKey(group string, host string) (string, error) {
	sshkey := inventory.All.Children[group].Hosts[host].SshKey
	if sshkey == "" {
		return "", errors.New("host sshkey does not exist")
	}
	return sshkey, nil
}

func (inventory *Inventory) getUngroupedHostSshKey(host string) (string, error) {
	sshkey := inventory.All.Hosts[host].SshKey
	if sshkey == "" {
		return "", errors.New("host sshkey does not exist")
	}
	return sshkey, nil
}
