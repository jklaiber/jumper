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

	password, err := s.GetHostPassword(groupName, hostName)
	if err != nil {
		return nil, fmt.Errorf("no valid host found")
	}
	sshKey, err := s.GetHostSSHKey(groupName, hostName)
	if err != nil {
		return nil, fmt.Errorf("no valid host found")
	}
	sshAgent, err := s.GetHostSSHAgent(groupName, hostName)
	if err != nil {
		return nil, fmt.Errorf("no valid host found")
	}
	sshAgentForwarding, err := s.GetHostSSHAgentForwarding(groupName, hostName)
	if err != nil {
		return nil, fmt.Errorf("no valid host found")
	}
	address, err := s.GetHostAddress(groupName, hostName)
	if err != nil {
		return nil, fmt.Errorf("no valid address found")
	}
	port, err := s.GetHostPort(groupName, hostName)
	if err != nil {
		return nil, fmt.Errorf("no valid host found")
	}

	return access.NewAccessConfig(username, password, address, port, sshKey, sshAgent, sshAgentForwarding), nil
}
