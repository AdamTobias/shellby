package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
		fmt.Print(text)
	}
}
