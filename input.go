package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

var ShellOperator = []string{"|", "<", ">"}

type Instruction struct {
	args []string
}
type Prompt struct {
	operators    []string
	instructions []Instruction
}

func GetInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	return text
}
func ParseInput(input string) (Prompt, error) {
	parts := strings.Fields(input)
	currentInstruction := Instruction{}
	prompt := Prompt{}
	if slices.Contains(ShellOperator, parts[0]) {
		return Prompt{}, fmt.Errorf("my_shell: syntax error near unexpected token `%s`", parts[0])
	}
	for _, word := range parts {
		if slices.Contains(ShellOperator, word) {
			prompt.instructions = append(prompt.instructions, currentInstruction)
			currentInstruction = Instruction{}
			prompt.operators = append(prompt.operators, word)
		} else {
			currentInstruction.args = append(currentInstruction.args, word)
		}
	}
	if len(currentInstruction.args) > 0 {

		prompt.instructions = append(prompt.instructions, currentInstruction)
	}
	return prompt, nil
}
