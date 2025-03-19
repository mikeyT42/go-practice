package main

import (
	"errors"
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
