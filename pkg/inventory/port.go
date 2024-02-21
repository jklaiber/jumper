package inventory

import "fmt"

func (s *InventoryService) GetHostPort(groupName, hostName string) (int, error) {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if port := getPortFromVars(hostVars); port != 22 {
				return port, nil
			}
			return getPortFromVars(s.Inventory.All.Vars), nil
		} else {
			return 0, fmt.Errorf("host not found")
		}
	}
	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if port := getPortFromVars(hostVars); port != 22 {
			return port, nil
		}
		if port := getPortFromVars(s.Inventory.All.Children[groupName].Vars); port != 22 {
			return port, nil
		}
		return getPortFromVars(s.Inventory.All.Vars), nil
	} else {
		return 0, fmt.Errorf("host not found")
	}
}

func getPortFromVars(vars Vars) int {
	if vars.Port == 0 {
		return 22
	}
	return vars.Port
}
