package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Genre int

const (
	Fiction Genre = iota
	NonFicton
	ScienceFiction
	Mystery
	Biography
)

type Book struct {
	title  string
	author string
	year   uint
	genre  Genre
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

	l := [...]Book{
		Book{
			title:  "The Great Gatsby",
			author: "F. Scott Fitzgerald",
			year:   1925,
			genre:  Fiction,
		},
		Book{
			title:  "Harry Potter and the Sorcer's Stone",
			author: "J. K. Rowling",
			year:   1997,
			genre:  Fiction,
		},
		Book{
			title:  "The War of the Worlds",
			author: "H. G. Wells",
			year:   1898,
			genre:  ScienceFiction,
		},
		Book{
			title:  "Sherlock Holmes",
			author: "Arthur Conan Doyle",
			year:   1892,
			genre:  Mystery,
		},
		Book{
			title:  "Steve Jobs",
			author: "Arthur Isaacson",
			year:   2011,
			genre:  Biography,
		},
		Book{
			title:  "Philosophiae Naturalis Principia Mathematica",
			author: "Sir Isaac Newton",
			year:   1687,
			genre:  NonFicton,
		},
	}

	fmt.Printf("Our library has %d book(s).\n\n", len(l))

	fmt.Print(
		"------------------------------------------------------------------\n",
		"                           Thank you\n",
		"------------------------------------------------------------------\n")
}
