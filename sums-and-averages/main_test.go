package main

import (
	"errors"
	"testing"
)

// -----------------------------------------------------------------------------
func TestValidateInputErrorArg(t *testing.T) {
	err := errors.New("I am a testing error")

	ve, e := validate("some input", err)
	if ve != nil {
		t.Errorf("actual %v :: expected %v", ve, nil)
	}
	_, ok := e.(*InputError)
	if !ok {
		t.Errorf("actual %T :: expected *InputError", e)
	}
}

// -----------------------------------------------------------------------------
func TestValidationNoInputError(t *testing.T) {
	ve, e := validate("\n", nil)
	if ve != nil {
		t.Errorf("actual %v :: expected %v", ve, nil)
	}
	_, ok := e.(*NoInputError)
	if !ok {
		t.Errorf("actual %T :: expected *NoInputError", e)
	}
}
