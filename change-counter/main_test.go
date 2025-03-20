package main

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// -----------------------------------------------------------------------------
func TestValidateInputErrorArg(t *testing.T) {
	const veEx float32 = 0
	err := errors.New("I am a testing error")

	ve, e := validate("some input", err)
	if ve != 0 {
		t.Errorf("actual %f :: expected %f", ve, veEx)
	}
	_, ok := e.(*InputError)
	if !ok {
		t.Errorf("actual %T :: expected *InputError", e)
	}
}

// -----------------------------------------------------------------------------
func TestValidationInputErrorParse(t *testing.T) {
	const veEx float32 = 0

	ve, e := validate("some input", nil)
	if ve != 0 {
		t.Errorf("actual %f :: expected %f", ve, veEx)
	}
	ie, ok := e.(*InputError)
	if !ok {
		t.Errorf("actual %T :: expected *InputError", e)
	}
	_, ok = ie.Err.(*strconv.NumError)
	if !ok {
		t.Errorf("actual %T :: expected *strconv.NumError", ie.Err)
	}
}

// -----------------------------------------------------------------------------
func TestValidationNoInputError(t *testing.T) {
	const veEx float32 = 0

	ve, e := validate("\n", nil)
	if ve != 0 {
		t.Errorf("actual %f :: expected %f", ve, veEx)
	}
	_, ok := e.(*NoInputError)
	if !ok {
		t.Errorf("actual %T :: expected *NoInputError", e)
	}
}

// -----------------------------------------------------------------------------
func TestValidationOutOfRangeError(t *testing.T) {
	const veEx float32 = 0

	ve, e := validate("5\n", nil)
	if ve != 0 {
		t.Errorf("actual %f :: expected %f", ve, veEx)
	}
	_, ok := e.(*OutOfRangeError)
	if !ok {
		t.Errorf("actual %T :: expected *InputError", e)
	}
}

// -----------------------------------------------------------------------------
func TestValidationSentinel(t *testing.T) {
	const exSentinel float32 = -1

	v, e := validate("-1\n", nil)
	if e != nil {
		t.Errorf("actual %T :: expected nil", e)
	}
	if v != exSentinel {
		t.Errorf("actual %f :: expected %f", v, exSentinel)
	}
}

// -----------------------------------------------------------------------------
func TestValidationSuccess(t *testing.T) {
	const exV float32 = 0.5

	v, e := validate("0.5\n", nil)
	if e != nil {
		t.Errorf("actual %T :: expected nil", e)
	}
	if v != exV {
		t.Errorf("actual %f :: expected %f", v, exV)
	}
}

// -----------------------------------------------------------------------------
func TestCalculateChange(t *testing.T) {
	var tests = []struct {
		cost                                                  float32
		exNumQuarters, exNumDimes, exNumNickels, exNumPennies int
	}{
		{0.50, 2, 0, 0, 0},
		{0.01, 0, 0, 0, 1},
		{0.99, 3, 2, 0, 4},
		{0.25, 1, 0, 0, 0},
		{0.45, 1, 2, 0, 0},
		{1.00, 4, 0, 0, 0},
		{0.00, 0, 0, 0, 0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("$%.2f, %d Q, %d D, %d N, %d P",
			tt.cost, tt.exNumQuarters, tt.exNumDimes, tt.exNumNickels,
			tt.exNumPennies)

		t.Run(testname, func(t *testing.T) {
			var numQuarters, numDimes, numNickels, numPennies int
			calculateChange(&tt.cost, &numQuarters, &numDimes, &numNickels,
				&numPennies)
			changeMatch := numQuarters != tt.exNumQuarters ||
				numDimes != tt.exNumDimes || numNickels != tt.exNumNickels ||
				numPennies != tt.exNumPennies

			if changeMatch {
				t.Errorf("actual [%d Q, %d D, %d N, %d P] :: "+
					"expected [%d Q, %d D, %d N, %d P]",
					numQuarters, numDimes, numNickels, numPennies,
					tt.exNumQuarters, tt.exNumDimes, tt.exNumNickels,
					tt.exNumPennies)
			}
		})
	}
}
