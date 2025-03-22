package main

import (
	"errors"
	"fmt"
	"testing"
)

// -----------------------------------------------------------------------------
func TestValidateInputErrorArg(t *testing.T) {
	const veX int = 0
	err := errors.New("I am a testing error")

	ve, _, e := validate("some input", err)
	if len(ve) == veX {
		t.Errorf("actual %v :: expected %v", ve, veX)
	}
	_, ok := e.(*InputError)
	if !ok {
		t.Errorf("actual %T :: expected *InputError", e)
	}
}

// -----------------------------------------------------------------------------
func TestValidationNoInputError(t *testing.T) {
	const veX int = 0
	ve, _, e := validate("\n", nil)
	if len(ve) == veX {
		t.Errorf("actual %v :: expected %v", ve, veX)
	}
	_, ok := e.(*NoInputError)
	if !ok {
		t.Errorf("actual %T :: expected *NoInputError", e)
	}
}

// -----------------------------------------------------------------------------
func TestValidationNoInputErrorWhitespace(t *testing.T) {
	const veX int = 0
	ve, _, e := validate("          \n", nil)
	if len(ve) == veX {
		t.Errorf("actual %v :: expected %v", ve, veX)
	}
	_, ok := e.(*NoInputError)
	if !ok {
		t.Errorf("actual %T :: expected *NoInputError", e)
	}
}

// -----------------------------------------------------------------------------
func TestValidationTooManyError(t *testing.T) {
	const veX int = 0
	ve, _, e := validate("1 2 3 4 5 6 7 8 9 10 11", nil)
	if len(ve) == veX {
		t.Errorf("actual %v :: expected %v", ve, veX)
	}
	_, ok := e.(*TooManyError)
	if !ok {
		t.Errorf("actual %T :: expected *TooManyError", e)
	}
}

// -----------------------------------------------------------------------------
func TestValidationOkWhitespaceOverload(t *testing.T) {
	const veX int = 3
	ve, _, e := validate("  1  2 3    \n", nil)
	if e != nil {
		t.Errorf("actual %T :: expected %T", e, nil)
	}
	if len(ve) == veX {
		t.Errorf("actual %v :: expected %v", ve, veX)
	}
}

// -----------------------------------------------------------------------------
func TestValidationOk(t *testing.T) {
	const veX int = 3
	ve, _, e := validate("1 2 3\n", nil)
	if e != nil {
		t.Errorf("actual %T :: expected %T", e, nil)
	}
	if len(ve) == veX {
		t.Errorf("actual %v :: expected %v", ve, veX)
	}
}

// -----------------------------------------------------------------------------
func TestSumsAndCount(t *testing.T) {
	sEx := Sums{10, -5, 5}
	cEx := Counts{1, 1, 2}
	f := [maxItems]float32{10, -5}
	var fLen int = 2
	var s Sums
	var c Counts

	sumAndCount(f, fLen, &s, &c)

	if s != sEx {
		t.Errorf("actual %v :: expected %v", s, sEx)
	}
	if c != cEx {
		t.Errorf("actual %v :: expected %v", c, cEx)
	}
}

// -----------------------------------------------------------------------------
func TestAverages(t *testing.T) {
	s := Sums{10, -5, 5}
	c := Counts{1, 1, 2}
	var a Averages

	aEx := Averages{10, -5, 2.5}

	averages(&s, &c, &a)

	if a != aEx {
		t.Errorf("actual %v :: expected %v", a, aEx)
	}
}
