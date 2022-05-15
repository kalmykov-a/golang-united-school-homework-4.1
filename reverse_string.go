package reverse_string

func ReverseString(input string) (output string) {
	// solution goes here
	for _, r := range input {
		output = string(r) + output
	}
	return output
}
