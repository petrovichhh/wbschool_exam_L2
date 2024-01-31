package main

import (
	"log"
	"testing"
)

func TestUnpackingString(t *testing.T) {
	testTable := []struct {
		input  string
		expect string
	}{
		{
			input:  "a4bc2d5e",
			expect: "aaaabccddddde",
		},
		{
			input:  "a3b3c3",
			expect: "aaabbbccc",
		},
		{
			input:  "a9b",
			expect: "aaaaaaaaab",
		},
		{
			input:  "abcd",
			expect: "abcd",
		},
		{
			input:  "",
			expect: "",
		},
	}

	for _, testCase := range testTable {
		output, err := unpackingString(testCase.input)
		if err != nil {
			log.Println(err)
		}
		if testCase.expect != output {
			t.Errorf("Incorrect result. Expect %s, got %s",
				testCase.expect,
				output)
		}
	}
}
