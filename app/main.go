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
	fmt.Print("$ ")

	command, err := bufio.NewReader(os.Stdin).ReadString('\n')

	if err != nil {
		print("Err: ", err)
		os.Exit(1)
	}

	fmt.Println(strings.TrimSpace(command), "Command not found:")

}
