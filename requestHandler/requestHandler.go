package requestHandler

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"github.com/pkg/sftp"
)

func HandleRequests(client *sftp.Client) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")	
		command, args := getCommandAndArgs(reader)
		fmt.Println(command)
		fmt.Println(args)
	}
}

func getCommandAndArgs(reader *bufio.Reader) (string, []string) {
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Unable to read input from user.")
		os.Exit(1)
	}

	tokens := strings.Split(input, " ")
	
	return tokens[0], tokens[1:]
}
