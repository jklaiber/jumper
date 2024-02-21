package inventory

import (
	"fmt"

	"github.com/jklaiber/jumper/pkg/access"
)

func (s *InventoryService) GetAccessConfig(groupName, hostName string) (*access.AccessConfig, error) {
	username, err := s.GetHostUsername(groupName, hostName)
	if err != nil {
		return nil, fmt.Errorf("no valid username found")
	}

	password := s.GetHostPassword(groupName, hostName)
	sshKey := s.GetHostSSHKey(groupName, hostName)
	sshAgent := s.GetHostSSHAgent(groupName, hostName)
	sshAgentForwarding := s.GetHostSSHAgentForwarding(groupName, hostName)
	address, err := s.GetHostAddress(groupName, hostName)
	if err != nil {
		return nil, fmt.Errorf("no valid address found")
	}
	port := s.GetHostPort(groupName, hostName)

	return access.NewAccessConfig(username, password, address, port, sshKey, sshAgent, sshAgentForwarding), nil
}
