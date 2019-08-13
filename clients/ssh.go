package clients

import (
	"os"
	"strings"
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"golang.org/x/crypto/ssh/terminal"
)

var username string = os.Getenv("USER")

func CreateSSHClient(addr string) (*ssh.Client, error) {
	parseAddrAndSetUser(addr)

	homeDir := os.Getenv("HOME")
	hostKeyCallback, err := knownhosts.New(homeDir + "/.ssh/known_hosts")

	if err != nil {
		return nil, err
	}

	config := generateClientConfig(hostKeyCallback)

	addrWithPort := addr + ":22"
	client, err := ssh.Dial("tcp", addrWithPort, config)

	return client, err
}

func parseAddrAndSetUser(addr string) {
	parts := strings.Split(addr, "@")

	if len(parts) > 1 {
		username = parts[0]
	}
}

func generateClientConfig(hostKeyCallback ssh.HostKeyCallback) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: username, 
		Auth: []ssh.AuthMethod{
			ssh.PasswordCallback(passwordPrompt),
		},
		HostKeyCallback: hostKeyCallback,
		Timeout: time.Duration(10) * time.Second,
	}
}

func passwordPrompt() (string, error) {
	var res []byte
	fmt.Print("Password for " + username + ": ")
	res, err := terminal.ReadPassword(0);

	return string(res), err
}
