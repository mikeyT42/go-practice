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

type InputValidationType int

const (
	Ok InputValidationType = iota
	OutOfRange
	NoInput
	InputError
	IOError
)

type InputValidation struct {
	Type  InputValidationType
	Value interface{}
}

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

	fmt.Print("Enter the amount you spent to two decimal places: the input\n",
		"must be between 0 and 1: -1 is to exit.\n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(sentinel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin %v\n", err)
		return Continue
	}
	input = input[:len(input)-1] // Get rid of the sentinel.

	return Continue
}

// -----------------------------------------------------------------------------
func validate(input string, err error) InputValidation {
	if err != nil {
		return InputValidation{Type: IOError, Value: err}
	}

	var f float32 = 0.0

	return InputValidation{Type: Ok, Value: f}
}
