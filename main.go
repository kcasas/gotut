package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")

		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		if text == "exit" {
			break
		}

		fmt.Println(text)
	}
}
