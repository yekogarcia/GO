package main

import "testing"

func TestSum(t *testing.T) {
	// result := Sum(2, 3)
	// expected := 5
	// if result != expected {
	// 	t.Errorf("Sum(2, 3) = %d; want %d", result, expected)
	// }

	tables := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{1, 1, 2},
		{0, 0, 0},
	}

	for _, tt := range tables {
		result := Sum(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a, b     int
		expected int
	}{
		{4, 3, 4},
		{3, 2, 3},
		{2, 5, 5},
	}

	for _, tt := range tables {
		max := getMax(tt.a, tt.b)
		if max != tt.expected {
			t.Errorf("getMax(%d, %d) = %d; want %d", tt.a, tt.b, max, tt.expected)
		}
	}
}
