package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)
			userCmd := cleanedInput[0]
			cmdFn, exists := commands[userCmd]
			if !exists {
				fmt.Println("Unknown command")
			} else {
				cmdFn.callback()

			}

		}

	}
}
