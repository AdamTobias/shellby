package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// what is os.Stdin
	reader := bufio.NewReader(os.Stdin)

	for {
		// why is it printing delim
		// why do I have to hit enter?
		fmt.Print("s: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("\nSee ya!")
			} else {
				fmt.Println("\nEncountered err: ", err)
			}
			return
		}

		cmdName, cmdArgs := parseInput(text)
		cmd := exec.Command(cmdName, cmdArgs...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			fmt.Print(cmdName + ": ")
			if errors.Is(err, exec.ErrNotFound) {
				fmt.Println("command not found")
			} else {
				fmt.Println(err)
			}
		}
	}
}

// parseInput cleans up the input and returns the command name and an array of arguments
func parseInput(input string) (string, []string) {
	input = strings.Trim(input, "\n")
	input = strings.Trim(input, " ")
	splitInput := strings.Split(input, " ")
	cmdName := splitInput[0]
	var allCmdArgs []string
	var cmdArgs []string

	if len(splitInput) > 1 {
		allCmdArgs = splitInput[1:]
	}
	// remove any empty arguments
	for _, arg := range allCmdArgs {
		if arg != "" {
			cmdArgs = append(cmdArgs, arg)
		}
	}
	return cmdName, cmdArgs
}
