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
		return nil, err
	}
	sshKey, err := s.GetHostSSHKey(groupName, hostName)
	if err != nil {
		return nil, err
	}
	sshAgent, err := s.GetHostSSHAgent(groupName, hostName)
	if err != nil {
		return nil, err
	}
	sshAgentForwarding, err := s.GetHostSSHAgentForwarding(groupName, hostName)
	if err != nil {
		return nil, err
	}
	address, err := s.GetHostAddress(groupName, hostName)
	if err != nil {
		return nil, err
	}
	port, err := s.GetHostPort(groupName, hostName)
	if err != nil {
		return nil, err
	}

	return access.NewAccessConfig(username, password, address, port, sshKey, sshAgent, sshAgentForwarding), nil
}
