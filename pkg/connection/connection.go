//go:generate mockgen -source=connection.go -destination=mocks/connection.go -package=mocks Connection
package connection

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jklaiber/jumper/pkg/access"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/term"
)

type Connection interface {
	Start()
	runShell(ctx context.Context) error
}

type SSHConnection struct {
	accessConfig access.Configurator
	sshConfig    *ssh.ClientConfig
}

func NewConnection(accessConfig access.Configurator) (*SSHConnection, error) {
	sshConfig, err := accessConfig.BuildClientConfig()
	if err != nil {
		return nil, fmt.Errorf("could not build SSH client config: %v", err)
	}

	return &SSHConnection{
		accessConfig: accessConfig,
		sshConfig:    sshConfig,
	}, nil
}

func (connection *SSHConnection) Start() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Printf("Connecting to %s with %s\n\n", connection.accessConfig.GetAddress(), connection.accessConfig.GetUsername())
	go func() {
		if err := connection.runShell(ctx); err != nil {
			log.Print(err)
		}
		cancel()
	}()

	select {
	case <-sig:
		cancel()
	case <-ctx.Done():
	}

}

func (connection *SSHConnection) runShell(ctx context.Context) error {
	if connection.accessConfig.GetPort() == 0 {
		connection.accessConfig.SetPort(22)
	}
	hostWithPort := fmt.Sprintf("%s:%d", connection.accessConfig.GetAddress(), connection.accessConfig.GetPort())
	conn, err := ssh.Dial("tcp", hostWithPort, connection.sshConfig)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Printf("could not close connection: %v", err)
		}
	}()

	session, err := conn.NewSession()
	if err != nil {
		return fmt.Errorf("cannot open new session: %v", err)
	}
	defer func() {
		err := session.Close()
		if err != nil {
			log.Printf("could not close session: %v", err)
		}
	}()

	if connection.accessConfig.GetSshAgentForwarding() {
		if err := agent.ForwardToRemote(conn, os.Getenv("SSH_AUTH_SOCK")); err != nil {
			return fmt.Errorf("error forwarding to remote: %v", err)
		}

		if err := agent.RequestAgentForwarding(session); err != nil {
			return fmt.Errorf("error requesting agent forwarding: %v", err)
		}
	}

	go func() {
		<-ctx.Done()
		err := conn.Close()
		if err != nil {
			log.Printf("could not close connection: %v", err)
		}
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

	// TODO check in ssh/sshd_config if enable ssh agent forwarding solves the problem...
	// When yes create a config option for the agent forwarding...

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
