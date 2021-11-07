package main

import "fmt"

func subset(input, output []int, index int) {
	if len(input) <= index {
		fmt.Println(output)
		return
	}
	subset(input, output, index+1)
	output = append(output, input[index])
	subset(input, output, index+1)
}
func main() {
	ARR := []int{1, 2, 3}
	output := []int{}
	subset(ARR, output, 0)
}
