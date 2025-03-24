package main

import "testing"

// -----------------------------------------------------------------------------
func TestCleanInputNoSpace(t *testing.T) {
	const (
		in   = "nursesrun\n"
		inEx = "nursesrun"
	)

	inAc := cleanInput(in)
	if inAc != inEx {
		t.Errorf("actual %s :: expected %s", inAc, inEx)
	}
}
