package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

var ShellOperator = []string{"|", "<", ">"}

type OperatorDequeue struct {
	top      int
	operator []string
}
type Instruction struct {
	args []string
}
type Prompt struct {
	operators          OperatorDequeue
	instructionDequeue []Instruction
}

func GetInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	return text
}
func ParseInput(input string) (Prompt, error) {
	opDequeue := OperatorDequeue{}
	currentInstruction := Instruction{}
	prompt := Prompt{}
	wordSlice := strings.Fields(input)
	if slices.Contains(ShellOperator, wordSlice[0]) {
		message := fmt.Errorf("my_shell : syntax error near unexpected token `%s`", wordSlice[0])
		return Prompt{}, message
	}
	for _, word := range wordSlice {
		if slices.Contains(ShellOperator, word) {
			prompt.instructionDequeue = append(prompt.instructionDequeue, in)
			currentInstruction = Instruction{}
			opDequeue.operator = append(opDequeue.operator, word)
		} else {
			currentInstruction.args = append(currentInstruction.args, word)
		}
	}
	if len(opDequeue.operator) > 0 {
		prompt.operators = opDequeue
	}
	if len(currentInstruction.args) > 0 {
		prompt.instructionDequeue = append(prompt.instructionDequeue, currentInstruction)
	}
	return prompt, nil
}
