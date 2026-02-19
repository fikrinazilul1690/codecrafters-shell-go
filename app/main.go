package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func parseInput(input string) (string, []string) {
	parts := strings.Fields(input)
	command, args := parts[0], parts[1:]
	return command, args
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
		path, err := exec.LookPath(command)
		if err != nil {
			fmt.Println(command + ": not found")
		}
		fmt.Println(command, "is", path)
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
			fmt.Println(command + ": command not found")
		}
	}
}
