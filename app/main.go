package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	for {
		fmt.Print("$ ")

		reader := bufio.NewReader(os.Stdin)

		command, err := reader.ReadString('\n')

		command = strings.TrimSpace(command)

		if err != nil {
			print("Err: ", err)
			os.Exit(1)
		}

		if command == "exit" {
			break
		}

		fmt.Printf("%s: command not found\n", command)

	}
}
