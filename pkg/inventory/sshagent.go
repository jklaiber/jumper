package inventory

import "errors"

func (inventory *Inventory) GetSshAgent(group string, host string) (bool, error) {
	if group != "" {
		sshagent, err := inventory.getGroupHostSshAgent(group, host)
		if err != nil {
			sshagent, err = inventory.getGroupSshAgent(group)
			if err != nil {
				sshagent, err = inventory.getGlobalSshAgent()
				if err != nil {
					return false, err
				}
				return sshagent, nil
			}
			return sshagent, nil
		}
		return sshagent, nil
	}
	sshagent, err := inventory.getUngroupedHostSshAgent(host)
	if err != nil {
		sshagent, err = inventory.getGlobalSshAgent()
		if err != nil {
			return false, err
		}
		return sshagent, nil
	}
	return sshagent, nil
}

func (inventory *Inventory) getGlobalSshAgent() (bool, error) {
	sshagent := inventory.All.Vars.SshAgent
	if !sshagent {
		return false, errors.New("global sshagent does not exist")
	}
	return sshagent, nil
}

func (inventory *Inventory) getGroupSshAgent(group string) (bool, error) {
	sshagent := inventory.All.Children[group].Vars.SshAgent
	if !sshagent {
		return false, errors.New("group sshagent does not exist")
	}
	return sshagent, nil
}

func (inventory *Inventory) getGroupHostSshAgent(group string, host string) (bool, error) {
	sshagent := inventory.All.Children[group].Hosts[host].SshAgent
	if !sshagent {
		return false, errors.New("host sshagent does not exist")
	}
	return sshagent, nil
}

func (inventory *Inventory) getUngroupedHostSshAgent(host string) (bool, error) {
	sshagent := inventory.All.Hosts[host].SshAgent
	if !sshagent {
		return false, errors.New("host sshagent does not exist")
	}
	return sshagent, nil
}
