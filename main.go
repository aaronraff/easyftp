package main

import (
	"os"
	"fmt"
	"errors"

	"github.com/aaronraff/easyftp/clients"
)

func main() {
	if err := validateArgs(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addr := os.Args[1]
	sshClient, err := clients.CreateSSHClient(addr)

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
