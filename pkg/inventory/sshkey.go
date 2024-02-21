package inventory

func (s *InventoryService) GetHostSSHKey(groupName, hostName string) string {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if sshkey := getSSHKeyFromVars(hostVars); sshkey != "" {
				return sshkey
			}
			return getSSHKeyFromVars(s.Inventory.All.Vars)
		}
	}

	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if sshkey := getSSHKeyFromVars(hostVars); sshkey != "" {
			return sshkey
		}
		if sshkey := getSSHKeyFromVars(s.Inventory.All.Children[groupName].Vars); sshkey != "" {
			return sshkey
		}
	}

	return getSSHKeyFromVars(s.Inventory.All.Vars)
}

func getSSHKeyFromVars(vars Vars) string {
	return vars.SshKey
}
