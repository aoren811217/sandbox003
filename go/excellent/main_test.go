package main
import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(20)
	if result != "Even" {
		t.Errorf("Expected 'Even', actual: %s", result)
	}
}
