package main

import (
	"os"
	"fmt"
	"errors"
	"strings"

	"github.com/aaronraff/easyftp/clients"
)

func main() {
	if err := validateArgs(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	addr := os.Args[1]
	user, host := parseAddrForUserAndHost(addr)
	_, err := clients.CreateSSHClient(host, user)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func validateArgs() error {
	if len(os.Args) < 2 {
		errMsg := fmt.Sprintf("Usage: %s address", os.Args[0])
		return errors.New(errMsg)
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
