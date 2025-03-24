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

	fmt.Print(
		"Please enter a string that is a palindrome; if you want to exit\n",
		"then just hit enter. It can be a sentence or a word.\n\n",
	)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(sentinel)
	if err != nil {
		fmt.Println("There was an error reading input. Please try again.")
		return Continue
	}

	input = cleanInput(input)

	return Continue
}

// -----------------------------------------------------------------------------
func cleanInput(s string) string {
	cl := strings.TrimSpace(s)

	return cl
}
