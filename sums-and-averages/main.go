package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type LoopControl int

const (
	Continue LoopControl = iota
	Stop
)

type NoInputError struct{}
type TooManyError struct{}
type InputError struct {
	Err error
}

func (e *NoInputError) Error() string {
	return "Sorry, but you didn't enter any input. Please try again."
}

func (e *TooManyError) Error() string {
	return "You entered too many numbers. Please try again."
}

func (e *InputError) Error() string {
	return fmt.Sprintf("You did not input valid input.\nerror:\n%v", e.Err)
}

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
	const sentinel = '\n'
	const maxItems = 10

	fmt.Printf(
		"Please input up to %d floating point or integer numbers. Seperate\n"+
			"them with spaces. Simply enter a newline character to exit.",
		maxItems)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(sentinel)

	validate(input, err)

	return Continue
}

// -----------------------------------------------------------------------------
func validate(input string, err error) ([]float32, error) {
	if err != nil {
		return nil, &InputError{err}
	}
	input = input[:len(input)-1]
	if len(input) == 0 {
		return nil, &NoInputError{}
	}

	for _, sn := range strings.Split(input, " ") {
	}

	var f []float32

	return f, nil
}
