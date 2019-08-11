package main

import (
	"os"
	"fmt"
	"errors"
)

func main() {
	if err := validateArgs(); err != nil {
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
