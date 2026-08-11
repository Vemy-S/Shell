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

		if command == " " {
			continue
		}

		args := parserQuoting(command)
		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		// strings.hasPrefix(command, "echo") found
		// strings.TrimPrefix(command, "echo") after
		// after, found := strings.CutPrefix(command, "echo")
		// if found {}
		if cmd == "exit" {
			break
		}

		//if after, found := strings.CutPrefix(command, "echo "); found {
		//fmt.Println(after)
		//continue
		//}
		//
		if cmd == "echo" {
			var cmdArgs = args[1:]
			fmt.Println(strings.Join(cmdArgs, " "))
			continue
		}

		// strings.hasPrefix(command, "type") - found
		// strings.TrimPrefix(command, "type") - target

		if cmd == "type" {
			if len(args) < 2 {
				continue
			}
			target := args[1]

			if targets[target] {
				fmt.Printf("%s is a shell builtin \n", target)
				continue
			}

			if !targets[target] {

				foundPath := findExecutable(target)

				if foundPath != "" {
					fmt.Printf("%s is %s\n", target, foundPath)
				} else {
					fmt.Printf("%s: not found\n", target)
				}
				continue
			}

		}

		foundPath := findExecutable(cmd)
		if foundPath != "" {
			executeProgram(cmd, args[1:])
			continue
		}

		fmt.Printf("%s: command not found\n", cmd)
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

func executeProgram(command string, args []string) {

	cmd := exec.Command(command, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()

	if err != nil {
		fmt.Printf("Err: %s \n Path: %s \n", err, command)
	}

}

func parserQuoting(command string) []string {
	var args []string
	var buf strings.Builder
	var inSingleQuotes = false

	for _, c := range command {
		if c == '\'' {
			inSingleQuotes = !inSingleQuotes
		} else if c == ' ' && !inSingleQuotes {
			if buf.Len() > 0 {
				args = append(args, buf.String())
				buf.Reset()
			}
		} else {
			buf.WriteRune(c)
		}
	}

	if buf.Len() > 0 {
		args = append(args, buf.String())
	}

	return args
}
