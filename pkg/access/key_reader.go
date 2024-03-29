package access

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func (a *AccessConfig) readSSHKeyPassphrase(file string) ([]byte, error) {
	fmt.Printf("Enter passphrase for key '%s': ", file)
	passphrase, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return nil, fmt.Errorf("error reading passphrase: %v", err)
	}
	return passphrase, nil
}

func (a *AccessConfig) parsePrivateKeyWithPassphrase(file string, buffer []byte) (ssh.AuthMethod, error) {
	passphrase, err := a.readSSHKeyPassphrase(file)
	if err != nil {
		return nil, fmt.Errorf("error reading the passphrase: %v", err)
	}
	key, err := ssh.ParsePrivateKeyWithPassphrase(buffer, passphrase)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}
	return ssh.PublicKeys(key), nil
}

func (a *AccessConfig) getPublicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading private key: %v", err)
	}
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		if _, ok := err.(*ssh.PassphraseMissingError); ok {
			return a.parsePrivateKeyWithPassphrase(file, buffer)
		}
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}
	return ssh.PublicKeys(key), nil
}
