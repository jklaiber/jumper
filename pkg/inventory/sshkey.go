package inventory

import "fmt"

func (s *InventoryService) GetHostSSHKey(groupName, hostName string) (string, error) {
	if groupName == "" {
		if hostVars, exists := s.Inventory.All.Hosts[hostName]; exists {
			if sshkey := getSSHKeyFromVars(hostVars); sshkey != "" {
				return sshkey, nil
			}
			return getSSHKeyFromVars(s.Inventory.All.Vars), nil
		} else {
			return "", fmt.Errorf("host not found")
		}
	}

	if hostVars, exists := s.Inventory.All.Children[groupName].Hosts[hostName]; exists {
		if sshkey := getSSHKeyFromVars(hostVars); sshkey != "" {
			return sshkey, nil
		}
		if sshkey := getSSHKeyFromVars(s.Inventory.All.Children[groupName].Vars); sshkey != "" {
			return sshkey, nil
		}
		return getSSHKeyFromVars(s.Inventory.All.Vars), nil
	} else {
		return "", fmt.Errorf("host not found")
	}
}

func getSSHKeyFromVars(vars Vars) string {
	return vars.SshKey
}
