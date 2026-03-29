package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		if len(parsedInput.operators) > 0 {
			HandleOperator(parsedInput)
		} else {
			cmd := exec.Command(parsedInput.instructions[0].args[0], parsedInput.instructions[0].args[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			}
		}
	}
}
