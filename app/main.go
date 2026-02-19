package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("$ ")
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("%v: command not found\n", command[:len(command)-1])
}
