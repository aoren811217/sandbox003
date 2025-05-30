package main

import "testing"
import "github.com/stretchr/testify/assert"


func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Even" {
		t.Errorf("Expected 'Even', actual: %s", result)
	}
}
