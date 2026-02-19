package main

import (
	"bufio"
	"fmt"
	"os"
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
		if _, ok := commands[args[0]]; ok {
			fmt.Println(args[0], "is a shell builtin")
			return
		}
		fmt.Println(args[0] + ": not found")
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
