package main

import (
	"fmt"
	"io"
	"os/exec"
	"slices"
)

var PIPE = []string{"|"}
var REDIRECTIONAL = []string{"<", ">"}

func handlePipe(pipe io.ReadCloser, prompt Prompt) {
	cmd_2 := exec.Command()
}
func handleRedirection() {

}
func HandleOpeator(prompt Prompt) {
	if slices.Contains(PIPE, prompt.operators.operator[prompt.operators.top]) {
		cmd_1 := exec.Command(prompt.instructionDequeue[0].args[0], prompt.instructionDequeue[0].args[1:]...)
		pipe, err := cmd_1.StdoutPipe()
		if err != nil {
			fmt.Println(err)
		}
		handlePipe(pipe, prompt.instructionDequeue[1:])
	} else if slices.Contains(REDIRECTIONAL, prompt.operators.operator[prompt.operators.top]) {

	}
}
