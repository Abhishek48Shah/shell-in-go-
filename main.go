package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("[my_shell@_<]: ")
		input := GetInput(scanner)
		if input == "" {
			continue
		}
		parsedInput, err := ParseInput(input)
		if err != nil {
			fmt.Println(err)
		}
		ProcessCommand(parsedInput, os.Stdin, os.Stdout)
	}
}
