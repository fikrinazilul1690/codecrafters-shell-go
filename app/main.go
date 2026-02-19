package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')
		sliceOfCommands := strings.Split(strings.TrimSpace(command), " ")
		command = sliceOfCommands[0]
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		if command == "echo" {
			fmt.Println(strings.Join(sliceOfCommands[1:], " "))
			continue
		}
		if command == "exit" {
			break
		}
		fmt.Printf("%v: command not found\n", command)
	}
}
