package main

import (
	"regexp"
	"fmt"
	"strings"
)

func NormalizeMobileNo(input []string) map[string]int {
	var output map[string]int
	output = make(map[string]int)

	for _,item:= range input {
		reg, err := regexp.Compile("[^0-9]+")
    	if err != nil {
       		fmt.Println(err)
		}
		fmt.Printf("Input : %v\n", item)
		result := reg.ReplaceAllString(item, "")
		fmt.Printf("Output : %v\n\n", result)		
		elem, isExist := output[result]
		if isExist {
			elem++
			output[result] = elem
		} else {
			output[result] = 1
		}
		
	}

	fmt.Println(output)
	return output
}

func NormalizeMobileNoNonRegexp(input []string) map[string]int {
	var output map[string]int
	output = make(map[string]int)

	for _,item:= range input {
		fmt.Printf("Input : %v\n", item)
		result := replaceSpecialChar(item)
		fmt.Printf("Output : %v\n\n", result)		
		elem, isExist := output[result]
		if isExist {
			elem++
			output[result] = elem
		} else {
			output[result] = 1
		}
	}
	fmt.Println(output)
	return output
}

func replaceSpecialChar(input string) string {
	input = strings.ReplaceAll(input," ","")
	input = strings.ReplaceAll(input,"(","")
	input = strings.ReplaceAll(input,")","")
	input = strings.ReplaceAll(input,"-","")
	return input
}

func main(){
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
	NormalizeMobileNo(input)
}