package connection

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/term"
)

func readSSHKeyPassphrase(file string) ([]byte, error) {
	fmt.Printf("Enter passphrase for key '%s': ", file)
	passphrase, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return nil, fmt.Errorf("error reading passphrase: %v", err)
	}
	return passphrase, nil
}

func parsePrivateKeyWithPassphrase(file string, buffer []byte) (ssh.AuthMethod, error) {
	passphrase, err := readSSHKeyPassphrase(file)
	if err != nil {
		return nil, fmt.Errorf("error reading the passphrase: %v", err)
	}
	key, err := ssh.ParsePrivateKeyWithPassphrase(buffer, passphrase)
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}
	return ssh.PublicKeys(key), nil
}

func PublicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading private key: %v", err)
	}
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		if _, ok := err.(*ssh.PassphraseMissingError); ok {
			return parsePrivateKeyWithPassphrase(file, buffer)
		}
		return nil, fmt.Errorf("error parsing private key: %v", err)
	}
	return ssh.PublicKeys(key), nil
}

func SSHAgent() (ssh.AuthMethod, error) {
	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, fmt.Errorf("error using SSH agent: %v", err)
	}
	return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers), nil
}

func NewConnection(username, host string, port int, password, sshkey string, sshagent bool) error {
	sshConfig := &ssh.ClientConfig{
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth:            []ssh.AuthMethod{},
	}

	if password != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.Password(password))
	}

	if sshkey != "" {
		parsedKey, err := PublicKeyFile(sshkey)
		if err != nil {
			return fmt.Errorf("could not read SSH key: %v", err)
		}
		sshConfig.Auth = append(sshConfig.Auth, parsedKey)
	}

	if sshagent {
		agentAuth, err := SSHAgent()
		if err != nil {
			log.Printf("SSH agent error: %v", err)
		} else {
			sshConfig.Auth = append(sshConfig.Auth, agentAuth)
		}
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Printf("Connecting to %s with %s\n\n", host, username)
	go func() {
		if err := runShell(ctx, host, port, sshConfig); err != nil {
			log.Print(err)
		}
		cancel()
	}()

	select {
	case <-sig:
		cancel()
	case <-ctx.Done():
	}

	return nil
}

func runShell(ctx context.Context, host string, port int, sshConfig *ssh.ClientConfig) error {
	if port == 0 {
		port = 22
	}
	hostWithPort := fmt.Sprintf("%s:%d", host, port)
	conn, err := ssh.Dial("tcp", hostWithPort, sshConfig)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return fmt.Errorf("cannot open new session: %v", err)
	}
	defer session.Close()

	go func() {
		<-ctx.Done()
		conn.Close()
	}()

	fd := int(os.Stdin.Fd())
	state, err := term.MakeRaw(fd)
	if err != nil {
		return fmt.Errorf("terminal make raw: %s", err)
	}

	defer func() {
		if err := term.Restore(fd, state); err != nil {
			log.Printf("could not restore terminal: %v", err)
		}
	}()

	w, h, err := term.GetSize(fd)
	if err != nil {
		return fmt.Errorf("terminal get size: %s", err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	term := os.Getenv("TERM")
	if term == "" {
		term = "xterm-256color"
	}
	if err := session.RequestPty(term, h, w, modes); err != nil {
		return fmt.Errorf("session xterm: %s", err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	if err := session.Shell(); err != nil {
		return fmt.Errorf("session shell: %s", err)
	}

	if err := session.Wait(); err != nil {
		if e, ok := err.(*ssh.ExitError); ok {
			switch e.ExitStatus() {
			case 130:
				return nil
			}
		}
		return fmt.Errorf("ssh: %s", err)
	}
	return nil
}
