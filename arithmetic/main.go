package main

import (
	"fmt"
	"math"
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
	fmt.Print(
		"------------------------------------------------------------------\n",
		"                           Welcome\n",
		"------------------------------------------------------------------\n")

	const x = 5
	const y = 2
	var result int

	result = add(x, y)
	fmt.Printf("result = %d\n\n", result)

	result = subtract(x, y)
	fmt.Printf("result = %d\n\n", result)

	result = multiply(x, y)
	fmt.Printf("result = %d\n\n", result)

	result = divide(x, y)
	fmt.Printf("result = %d\n\n", result)

	result = remainderOf(x, y)
	fmt.Printf("result = %d\n\n", result)

	result = power(x, y)
	fmt.Printf("result = %d\n\n", result)

	fmt.Print(
		"------------------------------------------------------------------\n",
		"                           Thank you\n",
		"------------------------------------------------------------------\n")
}

// -----------------------------------------------------------------------------
func add(x int, y int) int {
	fmt.Printf("adding %d to %d\n", x, y)
	return x + y
}

// -----------------------------------------------------------------------------
func subtract(x int, y int) int {
	fmt.Printf("subtracting %d from %d\n", y, x)
	return x - y
}

// -----------------------------------------------------------------------------
func multiply(x int, y int) int {
	fmt.Printf("multiplying %d by %d\n", x, y)
	return x * y
}

// -----------------------------------------------------------------------------
func divide(x int, y int) int {
	fmt.Printf("dividing %d by %d\n", x, y)
	return x / y
}

// -----------------------------------------------------------------------------
func remainderOf(x int, y int) int {
	fmt.Printf("remainder of the division of %d by %d\n", x, y)
	return x % y
}

// -----------------------------------------------------------------------------
func power(x int, y int) int {
	fmt.Printf("raising %d to the power of %d\n", x, y)
	return int(math.Pow(float64(x), float64(y)))
}
