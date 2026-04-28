package main

import (
	"bufio"
	"fmt"
	"slices"
	"strings"
)

var shellOperators = []string{"|", "<", ">", ">>", "&&"}
var redirectionOperators = []string{"<", ">", ">>"}

type Node struct {
	left  *Node
	args  []string
	right *Node
}
type Token struct {
	instruction []string
}

func GetInput(scanner *bufio.Scanner) string {
	scanner.Scan()
	text := scanner.Text()
	return text
}
func checkPrecedence(operator string) int {
	if strings.Compare(operator, "&&") == 0 || strings.Compare(operator, "||") == 0 {
		return 1
	} else if strings.Contains(operator, "|") {
		return 2
	} else if slices.Contains(redirectionOperators, operator) {
		return 3
	}
	return 0
}
func checkOperator(tokens []Token) int {
	index := 0
	prevPrecedence := 999
	for i, operator := range tokens {
		if slices.Contains(shellOperators, operator.instruction[0]) {
			currPrecedence := checkPrecedence(operator.instruction[0])
			if currPrecedence <= prevPrecedence {
				index = i
				prevPrecedence = currPrecedence
			}
		}
	}
	return index
}

func buildTree(tokens []Token) *Node {
	if len(tokens) == 0 {
		return nil
	}
	index := checkOperator(tokens)
	if !slices.Contains(shellOperators, tokens[index].instruction[0]) {
		return &Node{args: tokens[index].instruction}
	}
	root := &Node{args: tokens[index].instruction}
	root.left = buildTree(tokens[:index])
	root.right = buildTree(tokens[index+1:])
	return root
}
func buildToken(parts []string) []Token {
	var tokens []Token
	current := Token{}
	for _, word := range parts {
		if strings.Contains(word, "\"") {
			word = strings.Trim(word, "\"")
		} else if strings.Contains(word, "'") {
			word = strings.Trim(word, "'")
		}
		if slices.Contains(shellOperators, word) {
			tokens = append(tokens, current)
			tokens = append(tokens, Token{instruction: []string{word}})
			current = Token{}
		} else {
			current.instruction = append(current.instruction, word)
		}
	}
	if len(current.instruction) > 0 {
		tokens = append(tokens, current)
	}
	return tokens
}
func ParseInput(input string) (*Node, error) {
	parts := strings.Fields(input)
	if slices.Contains(shellOperators, parts[0]) {
		return &Node{}, fmt.Errorf("my_shell: syntax error near unexpected token `%s`", parts[0])
	}
	tokens := buildToken(parts)
	promptTree := buildTree(tokens)
	return promptTree, nil
}
