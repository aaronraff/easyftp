package clients

import (
	"os"
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"golang.org/x/crypto/ssh/terminal"
)

var user string

func CreateSSHClient(host string, username string) (*ssh.Client, error) {
	user = username
	homeDir := os.Getenv("HOME")
	hostKeyCallback, err := knownhosts.New(homeDir + "/.ssh/known_hosts")

	if err != nil {
		return nil, err
	}

	config := generateClientConfig(hostKeyCallback, user)

	addr := host + ":22"
	client, err := ssh.Dial("tcp", addr, config)

	return client, err
}

func generateClientConfig(hostKeyCallback ssh.HostKeyCallback, user string) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: user, 
		Auth: []ssh.AuthMethod{
			ssh.PasswordCallback(passwordPrompt),
		},
		HostKeyCallback: hostKeyCallback,
		Timeout: time.Duration(10) * time.Second,
	}
}

func passwordPrompt() (string, error) {
	var res []byte
	fmt.Print("Password for " + user + ": ")
	res, err := terminal.ReadPassword(0);

	return string(res), err
}
