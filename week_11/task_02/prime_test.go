package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"отрицательное", -1, false},
		{"ноль", 0, false},
		{"один", 1, false},
		{"два", 2, true},
		{"три", 3, true},
		{"четыре", 4, false},
		{"семь", 7, true},
		{"девять", 9, false},
		{"одиннадцать", 11, true},
		{"сто", 100, false},
		{"сто один", 101, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPrime(tt.input)
			if result != tt.expected {
				t.Errorf("IsPrime(%d) = %v; ожидалось %v", tt.input, result, tt.expected)
			}
		})
	}
}

