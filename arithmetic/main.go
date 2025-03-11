package main

import (
	"os"
	"os/exec"
)

func main() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Run()
    if err != nil {
        panic("Couldn't execute the clear command.")
    }
}
