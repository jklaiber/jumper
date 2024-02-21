package inventory

import "fmt"

func (s *InventoryService) GetHostAddress(groupName, hostName string) (string, error) {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			address, err := getAddressFromVars(hostVars)
			if err != nil {
				return "", err
			}
			return address, nil
		}
	}

	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		address, err := getAddressFromVars(hostVars)
		if err != nil {
			return "", err
		}
		return address, nil
	}

	return "", fmt.Errorf("no address found for host")
}

func getAddressFromVars(vars Vars) (string, error) {
	if vars.Address != "" {
		return vars.Address, nil
	}
	if vars.AnsibleHost != "" {
		return vars.AnsibleHost, nil
	}
	return "", fmt.Errorf("no address found for host")
}
