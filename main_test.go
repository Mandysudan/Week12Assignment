package main

import "testing"

// TestIsPositive tests the IsPositive function.
func TestIsPositive(t *testing.T) {
	if IsPositive(10) != true {
		t.Errorf("IsPositive(10) = %v; want true", IsPositive(10))
	}
	if IsPositive(-1) != false {
		t.Errorf("IsPositive(-1) = %v; want false", IsPositive(-1))
	}
	if IsPositive(0) != false {
		t.Errorf("IsPositive(0) = %v; want false", IsPositive(0))
	}
}
