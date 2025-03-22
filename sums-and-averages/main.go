package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const maxItems = 10

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

type Sums struct {
	positive float32
	negative float32
	overall  float32
}

type Counts struct {
	positive uint16
	negative uint16
	overall  uint16
}

type Averages struct {
	positive float32
	negative float32
	overall  float32
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

	fmt.Printf(
		"Please input up to %d floating point or integer numbers. Seperate\n"+
			"them with spaces. Simply enter a newline character to exit.\n\n",
		maxItems)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(sentinel)

	ns, nLen, err := validate(input, err)
	switch err := err.(type) {
	case *InputError:
		fmt.Printf(
			"You did not input valid input [%s]\nerror:\n%v\n\n",
			input[:len(input)-1], err.Err)
		return Continue
	case *TooManyError:
		fmt.Printf(
			"You input too many numbers, please only input up to %d.\n\n",
			maxItems,
		)
		return Continue
	case *NoInputError:
		return Stop
	}

	var s Sums
	var c Counts
	//var a Averages
	sumAndCount(ns, nLen, &s, &c)

	return Continue
}

// -----------------------------------------------------------------------------
func validate(input string, err error) ([maxItems]float32, int, error) {
	if err != nil {
		return [maxItems]float32{}, 0, &InputError{err}
	}
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return [maxItems]float32{}, 0, &NoInputError{}
	}

	iSplit := strings.Split(input, " ")
	if len(iSplit) == 0 {
		return [maxItems]float32{}, 0, &NoInputError{}
	}
	var f [maxItems]float32
	var fLen int = 0
	for i, sn := range iSplit {
		if i == maxItems {
			return [maxItems]float32{}, 0, &TooManyError{}
		}
		if sn == "" {
			continue
		}
		floc, err := strconv.ParseFloat(sn, 32)
		if err != nil {
			return [maxItems]float32{}, 0, &InputError{}
		}
		f[i] = float32(floc)
		fLen += 1
	}

	return f, 0, nil
}

// -----------------------------------------------------------------------------
func sumAndCount(ns [maxItems]float32, nLen int, s *Sums, c *Counts) {
	for i := range nLen {
		if ns[i] >= 0 {
			s.positive += ns[i]
			c.positive += 1
		} else {
			s.negative += ns[i]
			c.negative += 1
		}
		s.overall += ns[i]
		c.overall += 1
	}
}
