package main

import (
	"reflect"
	"testing"
)

func TestColumnToSort(t *testing.T) {
	data := [][]string{
		{"3", "apple", "banana"},
		{"1", "orange", "kiwi"},
		{"2", "grape", "mango"},
	}

	expected := [][]string{
		{"1", "orange", "kiwi"},
		{"2", "grape", "mango"},
		{"3", "apple", "banana"},
	}

	result := columnToSort(data, 1, true, false)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	data := [][]string{
		{"apple", "banana"},
		{"orange", "kiwi"},
		{"apple", "banana"},
	}

	expected := [][]string{
		{"apple", "banana"},
		{"orange", "kiwi"},
	}

	result := removeDuplicates(data)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
