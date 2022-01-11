package main

import "testing"

func TestSum(t *testing.T)  {
	t.Run("Is valid sum", func(t *testing.T) {
		sum := Sum(15, 15)
		expected := 30
	
		if sum != expected {
			t.Errorf("Invalid sum result: Result %d. Expected %d", sum, expected)
		}
	})
	t.Run("Is invalid sum", func(t *testing.T) {
		sum := Sum(14, 15)
		expected := 30
	
		if sum != expected {
			t.Errorf("Invalid sum result: Result %d. Expected %d", sum, expected)
		}
	})
}