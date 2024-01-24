// math_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	// Test case 1
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) returned %d, expected %d", result, expected)
	}

	// Test case 2
	result = Add(-1, 1)
	expected = 0
	if result != expected {
		t.Errorf("Add(-1, 1) returned %d, expected %d", result, expected)
	}

	// You can add more test cases as needed
}
