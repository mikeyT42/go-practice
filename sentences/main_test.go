package main

import "testing"

// -----------------------------------------------------------------------------
func TestKeystrokes(t *testing.T) {
	const (
		i  = "I am Michael Torres. I am 28."
		tk = 29
	)
	tkRes := keystrokes(i)
	if tkRes != tk {
		t.Errorf("actual %d :: expected %d", tkRes, tk)
	}
}

// -----------------------------------------------------------------------------
func TestAlphaRunes(t *testing.T) {
	const (
		i  = "I am Michael Torres. I am 28."
		ta = 19
	)
	taRes := alphaRunes(i)
	if taRes != ta {
		t.Errorf("actual %d :: expected %d", taRes, ta)
	}
}

// -----------------------------------------------------------------------------
func TestDigitRunes(t *testing.T) {
	const (
		i  = "I am Michael Torres. I am 28."
		td = 2
	)
	tdRes := digitRunes(i)
	if tdRes != td {
		t.Errorf("actual %d :: expected %d", tdRes, td)
	}
}

// -----------------------------------------------------------------------------
func TestVowelRunes(t *testing.T) {
	const (
		i  = "I am Michael Torres. I am 28."
		tv = 9
	)
	tvRes := vowelRunes(i)
	if tvRes != tv {
		t.Errorf("actual %d :: expected %d", tvRes, tv)
	}
}
