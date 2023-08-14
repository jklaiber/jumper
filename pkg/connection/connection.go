package connection

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/terminal"
)

func readSSHKeyPassphrase(file string) ([]byte, error) {
	fmt.Printf("Enter passphrase for key '%s':", file)
	passphrase, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()

	if err != nil {
		fmt.Printf("Error reading passphrase: %v\n", err)
		return nil, err
	}

	return passphrase, nil
}

func parsePrivateKeyWithPassphrase(file string, buffer []byte) (ssh.AuthMethod, error) {
	passphrase, err := readSSHKeyPassphrase(file)
	if err != nil {
		fmt.Errorf("Error reading the passphrase: %s", err)
		return nil, err
	}
	key, err := ssh.ParsePrivateKeyWithPassphrase(buffer, passphrase)
	if err != nil {
		return nil, fmt.Errorf("Error during parsing of pivate key; %v", err)
	}
	return ssh.PublicKeys(key), nil
}

func PublicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("Error during reading of private key; %v", err)
	}
	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		if _, ok := err.(*ssh.PassphraseMissingError); ok {
			return parsePrivateKeyWithPassphrase(file, buffer)
		} else {
			return nil, fmt.Errorf("Error during parsing of pivate key; %v", err)
		}
	}
	return ssh.PublicKeys(key), nil
}

func SSHAgent() ssh.AuthMethod {
	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	log.Fatal("There was an error using the ssh agent")
	return nil
}

func NewConnection(username string, host string, password string, sshkey string, sshagent bool) error {
	sshConfig := &ssh.ClientConfig{}

	if password != "" {
		sshConfig = &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}
	if sshkey != "" {
		parsedKey, err := PublicKeyFile(sshkey)
		if err != nil {
			return fmt.Errorf("Could not read ssh key: %v", err)
		}
		sshConfig = &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				parsedKey,
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}
	if sshagent {
		sshConfig = &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				SSHAgent(),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Printf("Connecting to %s with %s\n\n", host, username)
	go func() {
		if err := runShell(ctx, host, sshConfig); err != nil {
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

func runShell(ctx context.Context, host string, sshConfig *ssh.ClientConfig) error {
	conn, err := ssh.Dial("tcp", host+":22", sshConfig)
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
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		return fmt.Errorf("terminal make raw: %s", err)
	}
	defer terminal.Restore(fd, state)

	w, h, err := terminal.GetSize(fd)
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
