//go:generate mockgen -source=access.go -destination=mocks/access.go -package=mocks Configurator
package access

import "golang.org/x/crypto/ssh"

type Configurator interface {
	BuildClientConfig() (*ssh.ClientConfig, error)
	readSSHKeyPassphrase(file string) ([]byte, error)
	parsePrivateKeyWithPassphrase(file string, buffer []byte) (ssh.AuthMethod, error)
	getPublicKeyFile(file string) (ssh.AuthMethod, error)
	getSshAgent() (ssh.AuthMethod, error)
	GetUsername() string
	GetPassword() string
	GetAddress() string
	GetPort() int
	SetPort(int)
	GetSshKey() string
	GetSshAgent() bool
	GetSshAgentForwarding() bool
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
		parsedKey, err := a.getPublicKeyFile(a.SshKey)
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = append(sshConfig.Auth, parsedKey)
	}

	if a.SshAgent {
		agentAuth, err := a.getSshAgent()
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = append(sshConfig.Auth, agentAuth)
	}

	return sshConfig, nil
}

func (a *AccessConfig) GetUsername() string {
	return a.Username
}

func (a *AccessConfig) GetPassword() string {
	return a.Password
}

func (a *AccessConfig) GetAddress() string {
	return a.Address
}

func (a *AccessConfig) GetPort() int {
	return a.Port
}

func (a *AccessConfig) SetPort(port int) {
	a.Port = port
}

func (a *AccessConfig) GetSshKey() string {
	return a.SshKey
}

func (a *AccessConfig) GetSshAgent() bool {
	return a.SshAgent
}

func (a *AccessConfig) GetSshAgentForwarding() bool {
	return a.SshAgentForwarding
}
