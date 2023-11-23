package inventory

import "fmt"

func (inventory *Inventory) GetSshAgent(group string, host string) (bool, error) {
	if group != "" {
		if sshagent, err := inventory.getGroupHostSshAgent(group, host); err == nil {
			return sshagent, nil
		}
		if sshagent, err := inventory.getGroupSshAgent(group); err == nil {
			return sshagent, nil
		}
	}
	if sshagent, err := inventory.getUngroupedHostSshAgent(host); err == nil {
		return sshagent, nil
	}
	return inventory.getGlobalSshAgent()
}

func (inventory *Inventory) getGlobalSshAgent() (bool, error) {
	return getSshAgentFromVars(inventory.All.Vars)
}

func (inventory *Inventory) getGroupSshAgent(group string) (bool, error) {
	return getSshAgentFromVars(inventory.All.Children[group].Vars)
}

func (inventory *Inventory) getGroupHostSshAgent(group string, host string) (bool, error) {
	return getSshAgentFromVars(inventory.All.Children[group].Hosts[host])
}

func (inventory *Inventory) getUngroupedHostSshAgent(host string) (bool, error) {
	return getSshAgentFromVars(inventory.All.Hosts[host])
}

func getSshAgentFromVars(vars Vars) (bool, error) {
	if vars.SshAgent {
		return vars.SshAgent, nil
	}
	return false, fmt.Errorf("ssh agent not enabled")
}
