package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	a := 10
	b := 15
	expected := 25

	res := Sum(a, b)
	if res != expected {
		t.Errorf("Sum of %d and %d is expected to be %d but it is coming as %d", a, b, expected, res)
	}
}
