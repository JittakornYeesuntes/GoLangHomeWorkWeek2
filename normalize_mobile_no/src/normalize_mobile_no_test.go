package main

import (
	"testing"
	"reflect"
)

func TestNormalizeMobileNoRegexp(t *testing.T){
	input := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	expected := map[string]int{
		"1234567890":2,
		"1234567891":1,
		"1234567892":3,
		"1234567893":1,
		"1234567894":1,
	}

	output := NormalizeMobileNo(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expect % #v, but got % #v", expected, output)
	}
	
}

func TestNormalizeMobileNoReplaceString(t *testing.T){
	input := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	expected := map[string]int{
		"1234567890":2,
		"1234567891":1,
		"1234567892":3,
		"1234567893":1,
		"1234567894":1,
	}

	output := NormalizeMobileNo(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("expect % #v, but got % #v", expected, output)
	}
	
}