package random

import (
	"testing"
)

func TestChoice(t *testing.T) {
	// Test with a slice of strings
	strings := []string{"apple", "banana", "orange", "grape"}
	randomString := Choice(strings)
	if !contains(strings, randomString) {
		t.Errorf("Choice function returned an element not in the original slice")
	}

	// Test with a slice of integers
	integers := []int{1, 2, 3, 4, 5}
	randomNumber := Choice(integers)
	if !contains(integers, randomNumber) {
		t.Errorf("Choice function returned an element not in the original slice")
	}
}

// contains checks if a slice contains a specific element
func contains[T comparable](arr []T, elem T) bool {
	for _, e := range arr {
		if e == elem {
			return true
		}
	}
	return false
}
