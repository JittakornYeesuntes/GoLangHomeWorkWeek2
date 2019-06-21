package main

import "fmt"

func RomanConvert (input int) string {
	output := ""
	fmt.Printf("Input : %v ",input)
	for input > 0 {
		switch {
		case input >= 100 :
			output += "C"
			input -= 100

		case input >= 90 :
			output += "XC"
			input -= 90

		case input >= 50 :
			output += "L"
			input -= 50

		case input >= 40 :
			output += "XL"
			input -= 40

		case input >= 10 :
			output += "X"
			input -= 10

		case input >= 9 :
			output += "IX"
			input -= 9

		case input >= 5 :
			output += "V"
			input -= 5

		case input >= 4 :
			output += "IV"
			input -= 4

		case input >= 1 :
			output += "I"
			input --

		default :
			return "ERROR, NOT IN CASE"
		}
	}
	fmt.Printf("Output : %v \n",output)
	return output
}

func main() {
	input := 88
	fmt.Println(RomanConvert(input))
}