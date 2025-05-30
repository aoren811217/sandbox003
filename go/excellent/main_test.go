package main

import "testing"

int a = 10

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(10)
	if result != "Even" {
		t.Errorf("Expected 'Even', actual: %s", result)
	}
}
