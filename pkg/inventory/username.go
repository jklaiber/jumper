package inventory

import "fmt"

func (s *InventoryService) GetHostUsername(groupName, hostName string) (string, error) {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if username := getUsernameFromVars(hostVars); username != "" {
				return username, nil
			}
			username := getUsernameFromVars(s.Inventory.All.Vars)
			if username != "" {
				return username, nil
			}
			return "", fmt.Errorf("no username found for host")
		}
	}

	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if username := getUsernameFromVars(hostVars); username != "" {
			return username, nil
		}
		if username := getUsernameFromVars(s.Inventory.All.Children[groupName].Vars); username != "" {
			return username, nil
		}
	}

	username := getUsernameFromVars(s.Inventory.All.Vars)
	if username != "" {
		return username, nil
	}

	return "", fmt.Errorf("no username found for host")
}

func getUsernameFromVars(vars Vars) string {
	if vars.Username != "" {
		return vars.Username
	}
	if vars.AnsibleUser != "" {
		return vars.AnsibleUser
	}
	return ""
}
