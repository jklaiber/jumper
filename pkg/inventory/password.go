package inventory

func (s *InventoryService) GetHostPassword(groupName, hostName string) string {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if password := getPasswordFromVars(hostVars); password != "" {
				return password
			}
			return getPasswordFromVars(s.Inventory.All.Vars)
		}
	}
	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if password := getPasswordFromVars(hostVars); password != "" {
			return password
		}
		if password := getPasswordFromVars(s.Inventory.All.Children[groupName].Vars); password != "" {
			return password
		}
	}
	return getPasswordFromVars(s.Inventory.All.Vars)
}

func getPasswordFromVars(vars Vars) string {
	if vars.Password != "" {
		return vars.Password
	}
	if vars.AnsibleSshPASS != "" {
		return vars.AnsibleSshPASS
	}
	return ""
}
