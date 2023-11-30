package inventory

import "fmt"

type AccessConfig struct {
	Username           string
	Password           string
	Address            string
	Port               int
	SshKey             string
	SshAgent           bool
	SshAgentForwarding bool
}

func NewAccessConfig(username string, password string, address string, port int, sshKey string, sshAgent bool, sshAgentForwarding bool) *AccessConfig {
	return &AccessConfig{
		Username:           username,
		Password:           password,
		Address:            address,
		Port:               port,
		SshKey:             sshKey,
		SshAgent:           sshAgent,
		SshAgentForwarding: sshAgentForwarding,
	}
}

func (inventory *Inventory) GetAccessConfig(group string, host string) (*AccessConfig, error) {
	username, err := inventory.GetUsername(group, host)
	if err != nil {
		return nil, fmt.Errorf("no valid username found")
	}

	password, sshKey, sshAgent, sshAgentForwarding, err := inventory.getAccessMethod(group, host)
	if err != nil {
		return nil, fmt.Errorf("no valid access method found")
	}

	address, port, err := inventory.GetAddress(group, host)
	if err != nil {
		return nil, fmt.Errorf("no valid address found")
	}

	return NewAccessConfig(username, password, address, port, sshKey, sshAgent, sshAgentForwarding), nil
}

func (inventory *Inventory) getAccessMethod(group string, host string) (string, string, bool, bool, error) {
	password, passErr := inventory.GetPassword(group, host)
	sshkey, sshKeyErr := inventory.GetSshKey(group, host)
	sshAgent, sshAgentErr := inventory.GetSshAgent(group, host)
	sshAgentForwarding, sshAgentForwardingErr := inventory.GetSshAgentForwarding(group, host)

	if passErr == nil || sshKeyErr == nil || sshAgentErr == nil || sshAgentForwardingErr == nil {
		return password, sshkey, sshAgent, sshAgentForwarding, nil
	}

	return "", "", false, false, fmt.Errorf("no valid access method found")
}
