package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type LoopControl int

const (
	Continue LoopControl = iota
	Stop
)

type OutOfRangeError struct{}
type NoInputError struct{}
type InputError struct {
	Err error
}

func (e *OutOfRangeError) Error() string {
	return "The input was outside the range of 0 and 1."
}

func (e *NoInputError) Error() string {
	return "Sorry, but you didn't enter any input. Please try again."
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
	const delim = '\n'
	const sentinel = -1

	fmt.Print("Enter the amount you spent to two decimal places: the input\n",
		"must be between 0 and 1: -1 is to exit.\n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(delim)

	cost, err := validate(input, err)
	switch err := err.(type) {
	case *OutOfRangeError:
		fmt.Printf("You input %s, a value that is not between 0 and 1.\n\n",
			input[:len(input)-1])
		return Continue
	case *NoInputError:
		fmt.Print(
			"Sorry, but you didn't enter any input. Please try again.\n\n")
		return Continue
	case *InputError:
		fmt.Printf(
			"You did not input valid input [%s]\nerror:\n%v\n\n",
			input[:len(input)-1], err.Err)
		return Continue
	}
	if cost == sentinel {
		return Stop
	}

	var numQuarters int
	var numDimes int
	var numNickels int
	var numPennies int
	calculateChange(&cost, &numQuarters, &numDimes, &numNickels, &numPennies)
	fmt.Printf(
		"\nThe amount you gave was $%.2f, your change is %d Quarters,\n"+
			"%d Dimes, %d Nickels, and %d Pennies.\n\n", cost, numQuarters,
		numDimes, numNickels, numPennies,
	)

	return Continue
}

// -----------------------------------------------------------------------------
func validate(input string, err error) (float32, error) {
	const sentinel float32 = -1

	if err != nil {
		return 0, &InputError{err}
	}
	input = input[:len(input)-1] // Get rid of the sentinel.
	if len(input) == 0 {
		return 0, &NoInputError{}
	}

	f, err := strconv.ParseFloat(input, 32)
	if err != nil {
		return 0, &InputError{err}
	}

	fin := float32(f)
	if fin == sentinel {
		return fin, nil
	}
	if fin < 0 || fin > 1 {
		return 0, &OutOfRangeError{}
	}

	return fin, nil
}

// -----------------------------------------------------------------------------
func calculateChange(cost *float32, numQuarters *int, numDimes *int,
	numNickels *int, numPennies *int) {
	costCents := int(*cost * 100)
	change := costCents

	*numQuarters = change / 25
	change %= 25
	*numDimes = change / 10
	change %= 10
	*numNickels = change / 5
	change %= 5
	*numPennies = change
}
