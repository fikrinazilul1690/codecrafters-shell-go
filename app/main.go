package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
)

func parseInput(input string) (string, []string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return "", parts
	}
	command, args := parts[0], parts[1:]
	return command, args
}

func lookUpCommand(command string, fn func(path string)) error {
	path, err := exec.LookPath(command)
	if err != nil {
		return err
	}
	fn(path)
	return nil
}

func main() {
	commands := map[string]func([]string){
		"exit": func(args []string) { os.Exit(0) },
		"echo": func(args []string) { fmt.Println(strings.Join(args, " ")) },
	}
	commands["type"] = func(args []string) {
		command := args[0]
		if _, ok := commands[command]; ok {
			fmt.Println(command, "is a shell builtin")
			return
		}
		if err := lookUpCommand(command, func(path string) {
			fmt.Println(command, "is", path)
		}); err != nil {
			fmt.Println(command + ": not found")
		}
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		input = strings.TrimSpace(input)
		command, args := parseInput(input)
		if commandFunc, ok := commands[command]; ok {
			commandFunc(args)
		} else {
			if reflect.ValueOf(command).IsZero() {
				continue
			}
			if err := lookUpCommand(command, func(path string) {
				cmd := exec.Command(command, args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Run()
			}); err != nil {
				fmt.Println(command + ": command not found")
			}
		}
	}
}
