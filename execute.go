package main

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func executeCommand(node *Node, in io.Reader, out io.Writer) {
	cmd := exec.Command(node.args[0], node.args[1:]...)
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	}

}
func handlePipe(node *Node, in io.Reader, out io.Writer) {
	pr, pw := io.Pipe()
	go func() {
		ProcessCommand(node.left, in, pw)
		pw.Close()
	}()
	ProcessCommand(node.right, pr, out)

}
func handleRedirectionalOut(node *Node, in io.Reader) {
	if node.right == nil {
		return
	}
	file, err := os.Create(node.right.args[0])
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	defer file.Close()
	ProcessCommand(node.left, in, file)
}
func handleRedirectionalIn(node *Node, out io.Writer) {
	if node.right == nil {
		return
	}
	file, err := os.Open(node.right.args[0])
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	defer file.Close()
	ProcessCommand(node.left, file, out)
}
func handleRedirectionalAppend(node *Node, in io.Reader) {
	if node.right == nil {
		return
	}
	file, err := os.OpenFile(node.right.args[0], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	defer file.Close()
	ProcessCommand(node.left, in, file)
}
func handleCommandChanning(node *Node, in io.Reader, out io.Writer) {
	if node.right == nil {
		return
	}
	ProcessCommand(node.left, in, out)
	ProcessCommand(node.right, in, out)
}
func handleChangeDir(node *Node) {
	var targetDir string
	var err error
	if node == nil {
		return
	}
	sliceArr := node.args
	if len(sliceArr) == 1 {
		targetDir, err = os.UserHomeDir()
		targetDir += "/buffer/index"
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
		}
	} else {
		targetDir = strings.Trim(sliceArr[1], "[]")
	}
	err = os.Chdir(targetDir)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
	}
}
func ProcessCommand(node *Node, in io.Reader, out io.Writer) {
	if node == nil {
		return
	}
	switch node.args[0] {
	case "|":
		handlePipe(node, in, out)
	case ">":
		handleRedirectionalOut(node, in)
	case "<":
		handleRedirectionalIn(node, out)
	case ">>":
		handleRedirectionalAppend(node, in)
	case "&&":
		handleCommandChanning(node, in, out)
	case "exit":
		os.Exit(0)
	case "cd":
		handleChangeDir(node)
	default:
		executeCommand(node, in, out)
	}
}
