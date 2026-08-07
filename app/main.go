package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

var targets = map[string]bool{
	"type": true,
	"exit": true,
	"echo": true,
}

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
		}

		if after, found := strings.CutPrefix(command, "echo "); found {
			fmt.Println(after)
			continue
		}

		// strings.hasPrefix(command, "type") - found
		// strings.TrimPrefix(command, "type") - target

		if target, found := strings.CutPrefix(command, "type "); found {

			if targets[target] {
				fmt.Printf("%s is a shell builtin \n", target)
				continue
			}

			if !targets[target] {

				findExecutable(target)

				continue
			}

		}

		parts := strings.Fields(command)

		programFound := findExecutable(parts[0])

		if programFound != "" {
			executeProgram(programFound, parts[1:])
			continue
		}

		fmt.Printf("%s: command not found\n", command)
	}
}

func findExecutable(target string) string {

	path := os.Getenv("PATH")

	var foundPath string

	paths := filepath.SplitList(path)

	for _, path := range paths {

		fullpath := filepath.Join(path, target)
		info, err := os.Stat(fullpath)

		if err != nil {
			continue
		}

		isExecutable := info.Mode().Perm()&0110 != 0

		if isExecutable && !info.IsDir() {
			foundPath = fullpath
			break
		}

	}

	return foundPath
}

func executeProgram(programPath string, args []string) {

	cmd := exec.Command(programPath, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		fmt.Printf("Err: %s \n Path: %s \n", err, programPath)
	}

}
