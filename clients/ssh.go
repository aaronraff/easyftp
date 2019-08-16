package clients

import (
	"os"
	"fmt"
	"time"
	"io/ioutil"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"golang.org/x/crypto/ssh/terminal"
)

var user string
var homeDir string

func CreateSSHClient(host string, username string) (*ssh.Client, error) {
	user = username
	homeDir = os.Getenv("HOME")
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
			ssh.PublicKeysCallback(obtainPublicKey),
			ssh.PasswordCallback(passwordPrompt),
		},
		HostKeyCallback: hostKeyCallback,
		Timeout: time.Duration(10) * time.Second,
	}
}

func obtainPublicKey() ([]ssh.Signer, error) {
	key, err := ioutil.ReadFile(homeDir + "/.ssh/id_rsa")	
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	return []ssh.Signer { signer }, err
}

func passwordPrompt() (string, error) {
	var res []byte
	fmt.Print("Password for " + user + ": ")
	res, err := terminal.ReadPassword(0);

	return string(res), err
}
