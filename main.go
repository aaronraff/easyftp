package main

import (
	"os"
	"log"
	"errors"
	"strings"
	"sync"

	"github.com/aaronraff/easyftp/clients/ssh"
	"github.com/pkg/sftp"
)

func main() {
	if err := validateArgs(); err != nil {
		log.Fatal(err)
	}

	addr := os.Args[1]
	user, host := parseAddrForUserAndHost(addr)
	sshClient, err := ssh.CreateSSHClient(host, user)
	if err != nil {
		log.Fatal("Failed to create SSH client: ", err)
		os.Exit(1)
	}

	sftpClient, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatal("Failed to create SFTP session: ", err)
	}

	defer sftpClient.Close()
	startHandler(sftpClient)
}

func validateArgs() error {
	if len(os.Args) < 2 {
		return errors.New("Usage: " + os.Args[0] + " address")
	}

	return nil
}

func parseAddrForUserAndHost(addr string) (string, string) {
	parts := strings.Split(addr, "@")
	if len(parts) > 1 {
		return parts[0], parts[1]
	}

	return os.Getenv("USER"), parts[0]
}

func startHandler(sftpClient *sftp.Client) {
	var wg sync.WaitGroup
	wg.Add(1)
	go handleSftpRequests(sftpClient, wg)
	wg.Wait()
}
