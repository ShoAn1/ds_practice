package main

import "fmt"

func generate_parenthesis(outpur string, left, right int) []string {
	var output []string
	if left == 0 && right == 0 {
		output = append(output, outpur)
		return output
	} else {
		if left < right && left > 0 {
			output = append(output, generate_parenthesis(outpur+"(", left-1, right)...)
			output = append(output, generate_parenthesis(outpur+")", left, right-1)...)
			//fmt.Println(outpur)
		} else if left == 0 && right != 0 {
			output = append(output, generate_parenthesis(outpur+")", left, right-1)...)
			//	fmt.Println(outpur)
		} else {
			//	fmt.Println(outpur)
			output = append(output, generate_parenthesis(outpur+"(", left-1, right)...)
		}
		//fmt.Println(outpur)
	}
	return output
}

func main() {
	// arr := []string{}
	arr1 := generate_parenthesis("", 3, 3)
	fmt.Println(arr1)
}
