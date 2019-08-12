package main

import (
	"os"
	"fmt"
	"errors"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	if err := validateArgs(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addr := os.Args[1]
	sshClient, err := createSSHClient(addr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("test...")
	fmt.Println(sshClient)
}

func validateArgs() error {
	if len(os.Args) < 2 {
		errMsg := fmt.Sprintf("Usage: %s address", os.Args[0])
		return errors.New(errMsg)
	}

	return nil
}

func createSSHClient(addr string) (*ssh.Client, error) {
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

	client, err := ssh.Dial("tcp", addr, config)
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
