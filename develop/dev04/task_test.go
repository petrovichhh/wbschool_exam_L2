package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestAnagram(t *testing.T) {
	testTable := []struct {
		name     string
		input    []string
		expected map[string][]string
	}{
		{
			name:  "Test 1. Check for addition",
			input: []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{
				"акптя":  {"пятак", "пятка", "тяпка"},
				"иклост": {"листок", "слиток", "столик"},
			},
		},
		{
			name:     "Test 2. Less than two",
			input:    []string{"тяпка", "листок"},
			expected: map[string][]string{},
		},
		{
			name:  "Test 3. Dublicate",
			input: []string{"тяпка", "пятка", "пятка"},
			expected: map[string][]string{
				"акптя": {"пятка", "тяпка"},
			},
		},
	}
	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			actual, _ := anagram(tt.input)
			if !reflect.DeepEqual(tt.expected, actual) {
				t.Errorf("In %s, expected: %v, got: %v", tt.name, tt.expected, actual)
			}
		})
	}
}

func TestRemoveDublicate(t *testing.T) {
	input := []string{"тяпка", "пятка", "пятка"}
	expect := []string{"тяпка", "пятка"}

	result := removeDublicate(input)

	sort.Strings(result)
	sort.Strings(expect)

	if !reflect.DeepEqual(expect, result) {
		t.Errorf("expected: %v, got: %v", expect, result)
	}
}
