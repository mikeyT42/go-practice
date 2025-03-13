package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

type LoopControl int

const (
	Continue LoopControl = iota
	Stop
)

func main() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic("Couldn't execute the clear command.")
	}
	fmt.Print(
		"------------------------------------------------------------------\n",
		"                           Welcome\n",
		"------------------------------------------------------------------\n")

	var lc LoopControl = Continue
	for lc == Continue {
		lc = inputLoop()
	}

	fmt.Print(
		"------------------------------------------------------------------\n",
		"                           Thank you\n",
		"------------------------------------------------------------------\n")
}

// -----------------------------------------------------------------------------
func inputLoop() LoopControl {
	const sentinel = '\n'

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(sentinel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin %v\n", err)
		return Continue
	}
	if input[0] == sentinel {
		return Stop
	}

	return Continue
}
