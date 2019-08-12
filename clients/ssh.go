package clients

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func CreateSSHClient(addr string) (*ssh.Client, error) {
	var hostKey ssh.PublicKey
	username, password, err := getCredentials()

	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.FixedHostKey(hostKey),
		Timeout: time.Duration(10) * time.Second,
	}

	addrWithPort := fmt.Sprintf("%s:22", addr)
	client, err := ssh.Dial("tcp", addrWithPort, config)
	fmt.Println(hostKey)

	return client, err
}

func getCredentials() (string, string, error) {
	var username string
	fmt.Print("Username: ")
	fmt.Scanln(&username)

	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(0)

	if err != nil {
		return "", "", err
	}

	return username, string(password), nil
}
