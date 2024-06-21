package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		cliInput := reader.Text()
		clearInput(&cliInput)

		err := executeCommand(cliInput)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func clearInput(input *string) {
	*input = strings.TrimSpace(*input)
}
