package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// ehtan

func main() {
	// what is os.Stdin
	reader := bufio.NewReader(os.Stdin)

	for {
		// why is it printing delim
		// why do I have to hit enter?
		fmt.Print("s: ")
		text, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print("\nSee ya!\n")
			return
		}
		fmt.Print(text, "blah", err)
	}
}
