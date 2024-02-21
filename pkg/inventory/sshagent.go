package inventory

func (s *InventoryService) GetHostSSHAgent(groupName, hostName string) bool {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if sshAgent := getSSHAgentFromVars(hostVars); sshAgent {
				return true
			}
			return getSSHAgentFromVars(s.Inventory.All.Vars)
		}
	}
	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if sshAgent := getSSHAgentFromVars(hostVars); sshAgent {
			return true
		}
		if sshAgent := getSSHAgentFromVars(s.Inventory.All.Children[groupName].Vars); sshAgent {
			return true
		}
	}
	return getSSHAgentFromVars(s.Inventory.All.Vars)
}

func getSSHAgentFromVars(vars Vars) bool {
	return vars.SshAgent
}

func (s *InventoryService) GetHostSSHAgentForwarding(groupName, hostName string) bool {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if sshAgentForwarding := getSSHAgentForwardingFromVars(hostVars); sshAgentForwarding {
				return true
			}
			return getSSHAgentForwardingFromVars(s.Inventory.All.Vars)
		}
	}
	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if sshAgentForwarding := getSSHAgentForwardingFromVars(hostVars); sshAgentForwarding {
			return true
		}
		if sshAgentForwarding := getSSHAgentForwardingFromVars(s.Inventory.All.Children[groupName].Vars); sshAgentForwarding {
			return true
		}
	}
	return getSSHAgentForwardingFromVars(s.Inventory.All.Vars)
}

func getSSHAgentForwardingFromVars(vars Vars) bool {
	return vars.SshAgentForwarding
}
