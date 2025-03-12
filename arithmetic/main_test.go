package main

import "testing"

// -----------------------------------------------------------------------------
func TestAdd(t *testing.T) {
	const (
		x      = 2
		y      = 2
		result = 4
	)
	if add(x, y) != result {
		t.Errorf("add(%d, %d) is not equal to %d\n", x, y, result)
	}
}

// -----------------------------------------------------------------------------
func TestSubtract(t *testing.T) {
	const (
		x      = 2
		y      = 4
		result = -2
	)
	if subtract(x, y) != result {
		t.Errorf("subtract(%d, %d) is not equal to %d\n", x, y, result)
	}
}

// -----------------------------------------------------------------------------
func TestMultiply(t *testing.T) {
	const (
		x      = 2
		y      = 5
		result = 10
	)
	if multiply(x, y) != result {
		t.Errorf("multiply(%d, %d) is not equal to %d\n", x, y, result)
	}
}

// -----------------------------------------------------------------------------
func TestDivide(t *testing.T) {
	const (
		x      = 5
		y      = 2
		result = 2
	)
	if divide(x, y) != result {
		t.Errorf("divide(%d, %d) is not equal to %d\n", x, y, result)
	}
}

// -----------------------------------------------------------------------------
func TestRemainderOf(t *testing.T) {
	const (
		x      = 5
		y      = 2
		result = 1
	)
	if remainderOf(x, y) != result {
		t.Errorf("remainderOf(%d, %d) is not equal to %d\n", x, y, result)
	}
}

// -----------------------------------------------------------------------------
func TestPower(t *testing.T) {
	const (
		x      = 5
		y      = 2
		result = 25
	)
	if power(x, y) != result {
		t.Errorf("power(%d, %d) is not equal to %d\n", x, y, result)
	}
}
