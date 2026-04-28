package main

import (
	"bufio"
	"fmt"
	"os"
)

//	func printNode(node *Node) {
//		if node == nil {
//			return
//		}
//		printNode(node.left)
//		printNode(node.right)
//		arr := node.args
//		fmt.Println(strings.Trim(arr[1], "[]"))
//	}
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
