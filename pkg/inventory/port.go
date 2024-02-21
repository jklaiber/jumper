package inventory

func (s *InventoryService) GetHostPort(groupName, hostName string) int {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if port := getPortFromVars(hostVars); port != 22 {
				return port
			}
			return getPortFromVars(s.Inventory.All.Vars)
		}
	}
	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if port := getPortFromVars(hostVars); port != 22 {
			return port
		}
		if port := getPortFromVars(s.Inventory.All.Children[groupName].Vars); port != 22 {
			return port
		}
	}
	return getPortFromVars(s.Inventory.All.Vars)
}

func getPortFromVars(vars Vars) int {
	if vars.Port != 22 {
		return vars.Port
	}
	return 22
}
