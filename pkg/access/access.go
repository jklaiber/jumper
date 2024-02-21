package access

import "golang.org/x/crypto/ssh"

type Configurator interface {
	BuildClientConfig() (*ssh.ClientConfig, error)
}

type AccessConfig struct {
	Username           string
	Password           string
	Address            string
	Port               int
	SshKey             string
	SshAgent           bool
	SshAgentForwarding bool
}

var _ Configurator = (*AccessConfig)(nil)

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

func (a *AccessConfig) BuildClientConfig() (*ssh.ClientConfig, error) {
	sshConfig := &ssh.ClientConfig{
		User:            a.Username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{},
	}

	if a.Password != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.Password(a.Password))
	}

	if a.SshKey != "" {
		parsedKey, err := publicKeyFile(a.SshKey)
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = append(sshConfig.Auth, parsedKey)
	}

	if a.SshAgent {
		agentAuth, err := sshAgent()
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = append(sshConfig.Auth, agentAuth)
	}

	return sshConfig, nil
}
