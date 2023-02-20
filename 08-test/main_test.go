package main

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Even", 2, true},
		{"Odd", 3, false},
		{"Negative Even", -4, true},
		{"Negative Odd", -3, false},
		{"Zero", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isEven(tt.input)
			if got != tt.expected {
				t.Errorf("isEven(%d) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func isEven(num int) bool {
	return num%2 == 0
}
