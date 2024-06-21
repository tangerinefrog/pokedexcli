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

		cliRaw := reader.Text()
		if len(cliRaw) == 0 {
			continue
		}
		cliWords := getInput(cliRaw)
		c := cliWords[0]
		p := ""
		if len(cliWords) > 1 {
			p = cliWords[1]
		}

		err := executeCommand(c, p)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getInput(input string) []string {
	input = strings.TrimSpace(input)
	words := strings.Fields(input)
	return words
}
