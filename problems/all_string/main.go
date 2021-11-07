package main

import "fmt"

func all_ints(input, output []int, index int) {
	if index >= len(input) {
		// for _, i := range output {
		// }
		fmt.Println(output)
		return
	}
	all_ints(input, output, index+1)
	output = append(output, input[index])
	all_ints(input, output, index+1)
}

func all_string(input, output []string, index int) {
	if index >= len(input) {
		// for _, i := range output {
		// }
		fmt.Println(output)
		return
	}
	all_string(input, output, index+1)
	output = append(output, input[index])
	all_string(input, output, index+1)
}

func main() {
	// s1 := []int{1, 2, 3}
	// s2 := []int{}
	s23 := []string{"a", "b", "c"}
	s34 := []string{}
	all_string(s23, s34, 0)

	fmt.Println(s34)
}
