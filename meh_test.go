package main

import "testing"

func TestMehCanAddNumbers(t *testing.T) {
	sum := Meh(1, 2)

	if sum != 3 {
		t.Errorf("Incorrect sum; got %d, expected: %d", sum, 3)
	}

}
