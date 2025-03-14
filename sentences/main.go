package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"unicode"
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

	fmt.Printf(
		"\n\nPlease input a sentence. If you want to exit, just hit enter.\n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString(sentinel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading from stdin %v\n", err)
		return Continue
	}
	if input[0] == sentinel {
		return Stop
	}
	input = input[:len(input)-1] // Get rid of the sentinel.

	tk := keystrokes(input)
	fmt.Printf("\nKeystrokes: %10d", tk)
	ta := alphaRunes(input)
	fmt.Printf("\nAlpha Characters: %4d", ta)
	td := digitRunes(input)
	fmt.Printf("\nNumeric Characters: %2d", td)
	tv := vowelRunes(input)
	fmt.Printf("\nVowels: %14d", tv)

	return Continue
}

// -----------------------------------------------------------------------------
func keystrokes(input string) int {
	var k int = 0

	for range input {
		k += 1
	}

	return k
}

// -----------------------------------------------------------------------------
func alphaRunes(input string) int {
	var a int = 0

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}

		a += 1
	}

	return a
}

// -----------------------------------------------------------------------------
func digitRunes(input string) int {
	var d int = 0

	for _, r := range input {
		if !unicode.IsNumber(r) {
			continue
		}

		d += 1
	}

	return d
}

// -----------------------------------------------------------------------------
func vowelRunes(input string) int {
	var v int = 0

	for _, r := range input {
		switch r {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			v += 1
		}
	}

	return v
}
