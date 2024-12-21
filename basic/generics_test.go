package main

import (
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	tests := []struct {
		name      string
		input     []int
		expected1 []int
		expected2 []int
	}{
		{
			name:      "Even number of elements",
			input:     []int{1, 2, 3, 4},
			expected1: []int{1, 2},
			expected2: []int{3, 4},
		},
		{
			name:      "Odd number of elements",
			input:     []int{1, 2, 3, 4, 5},
			expected1: []int{1, 2},
			expected2: []int{3, 4, 5},
		},
		{
			name:      "Single element",
			input:     []int{1},
			expected1: []int{},
			expected2: []int{1},
		},
		{
			name:      "Empty slice",
			input:     []int{},
			expected1: []int{},
			expected2: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result1, result2 := SplitSlice(tt.input)
			if !reflect.DeepEqual(result1, tt.expected1) {
				t.Errorf("For %v, got first slice %v, expected %v", tt.input, result1, tt.expected1)
			}
			if !reflect.DeepEqual(result2, tt.expected2) {
				t.Errorf("For %v, got second slice %v, expected %v", tt.input, result2, tt.expected2)
			}
		})
	}
}
