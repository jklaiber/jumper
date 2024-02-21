package inventory

import "fmt"

func (s *InventoryService) GetHostSSHAgent(groupName, hostName string) (bool, error) {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if sshAgent := getSSHAgentFromVars(hostVars); sshAgent {
				return true, nil
			}
			return getSSHAgentFromVars(s.Inventory.All.Vars), nil
		} else {
			return false, fmt.Errorf("host not found")
		}
	}
	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if sshAgent := getSSHAgentFromVars(hostVars); sshAgent {
			return true, nil
		}
		if sshAgent := getSSHAgentFromVars(s.Inventory.All.Children[groupName].Vars); sshAgent {
			return true, nil
		}
		return getSSHAgentFromVars(s.Inventory.All.Vars), nil
	} else {
		return false, fmt.Errorf("host not found")
	}
}

func getSSHAgentFromVars(vars Vars) bool {
	return vars.SshAgent
}

func (s *InventoryService) GetHostSSHAgentForwarding(groupName, hostName string) (bool, error) {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if sshAgentForwarding := getSSHAgentForwardingFromVars(hostVars); sshAgentForwarding {
				return true, nil
			}
			return getSSHAgentForwardingFromVars(s.Inventory.All.Vars), nil
		} else {
			return false, fmt.Errorf("host not found")
		}
	}
	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if sshAgentForwarding := getSSHAgentForwardingFromVars(hostVars); sshAgentForwarding {
			return true, nil
		}
		if sshAgentForwarding := getSSHAgentForwardingFromVars(s.Inventory.All.Children[groupName].Vars); sshAgentForwarding {
			return true, nil
		}
		return getSSHAgentForwardingFromVars(s.Inventory.All.Vars), nil
	} else {
		return false, fmt.Errorf("host not found")
	}
}

func getSSHAgentForwardingFromVars(vars Vars) bool {
	return vars.SshAgentForwarding
}
