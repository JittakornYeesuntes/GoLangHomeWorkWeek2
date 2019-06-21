package main

import "testing"

func TestRomanConvert(t *testing.T){
	allTestCase := []struct{
		title string
		input int
		expected string
	}{
		{
			title : "Input 56 should return LVI",
			input : 56,
			expected : "LVI",
		},
		{
			title : "Input 74 should return LXXIV",
			input : 74,
			expected : "LXXIV",
		},
		{
			title : "Input 99 should return XCIX",
			input : 99,
			expected : "XCIX",
		},
		{
			title : "Input 48 should return XLVIII",
			input : 48,
			expected : "XLVIII",
		},

	}
	
	for _, testCase:= range allTestCase {
		t.Run(testCase.title, func (t *testing.T){
			output := RomanConvert(testCase.input)
			if testCase.expected != output {
				t.Errorf("expect % #v, but got % #v", testCase.expected, output)
			}
		})
	}
}