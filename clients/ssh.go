package clients

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func CreateSSHClient(addr string) (*ssh.Client, error) {
	certChecker := &ssh.CertChecker{}

	config := &ssh.ClientConfig{
		Auth: []ssh.AuthMethod{
			ssh.KeyboardInteractive(keyboardChallenge),
		},
		HostKeyCallback: certChecker.CheckHostKey,
		Timeout: time.Duration(10) * time.Second,
	}

	addrWithPort := fmt.Sprintf("%s:22", addr)
	client, err := ssh.Dial("tcp", addrWithPort, config)

	return client, err
}

func keyboardChallenge(user, instruction string, questions []string, echos []bool) (answers []string, err error) {
	var res []byte

	for i := 0; i < len(questions); i++ {
		fmt.Println(questions[i])

		if echos[i] == true {
			res, err = terminal.ReadPassword(0);
			answers[i] = string(res)
		} else {
			fmt.Scanln(&answers[i])
		}
	}

	return
}
