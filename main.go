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
		fmt.Print("[my_shell2 _<]: ")
		input := GetInput(scanner)
		if input == "" {
			continue
		}
		parsedInput, err := ParseInput(input)
		if err != nil {
			fmt.Print(err)
			continue
		}
		if len(parsedInput.operators.operator) > 0 {
			HandleOpeator(parsedInput)
		} else {
			cmd := exec.Command(parsedInput.instructionDequeue[0].args[0], parsedInput.instructionDequeue[0].args[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			if err := cmd.Run(); err != nil {
				fmt.Println(err)
			}
		}
	}
}
