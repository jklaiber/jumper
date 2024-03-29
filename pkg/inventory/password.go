package inventory

import "fmt"

func (s *InventoryService) GetHostPassword(groupName, hostName string) (string, error) {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if password := getPasswordFromVars(hostVars); password != "" {
				return password, nil
			}
			return getPasswordFromVars(s.Inventory.All.Vars), nil
		} else {
			return "", fmt.Errorf("host not found")
		}
	}
	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if password := getPasswordFromVars(hostVars); password != "" {
			return password, nil
		}
		if password := getPasswordFromVars(s.Inventory.All.Children[groupName].Vars); password != "" {
			return password, nil
		}
		return getPasswordFromVars(s.Inventory.All.Vars), nil
	} else {
		return "", fmt.Errorf("host not found")
	}
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
