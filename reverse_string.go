package reverse_string

import "unicode/utf8"

func ReverseString(input string) (output string) {
	// solution goes here
	var res []rune
	for i, r := range input {
		res[utf8.RuneCountInString(input)-1-i] = r
	}
	output = string(res)
	return output
}
