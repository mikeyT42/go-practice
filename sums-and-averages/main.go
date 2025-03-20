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

// -----------------------------------------------------------------------------
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
	const delim = '\n'

	fmt.Print("Enter the amount you spent to two decimal places: the input\n",
		"must be between 0 and 1: -1 is to exit.\n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(delim)

	return Continue
}
