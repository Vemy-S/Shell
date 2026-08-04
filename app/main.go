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
			os.Exit(0)
		}

		// strings.hasPrefix(command, "echo") found
		// strings.TrimPrefix(command, "echo") after
		// after, found := strings.CutPrefix(command, "echo")
		// if found {}
		if command == "exit" {
			break
		} else if after, found := strings.CutPrefix(command, "echo "); found {
			fmt.Println(after)
			continue
		}

		fmt.Printf("%s: command not found\n", command)

	}
}
