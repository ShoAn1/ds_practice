package main

import "fmt"

func check_palindrome(input int) int {
	if input == 0 {
		// fmt.Println("0")
		return 1
	}
	if input < 0 {
		// fmt.Println("<0 ", input)
		return 0
	}
	// fmt.Println("value is ", input)
	return check_palindrome(input-1) + check_palindrome(input-2) + check_palindrome(input-3)

}
func main() {
	//arr := []int{1, 2, 3, 5, 7}
	fmt.Println(check_palindrome(212))
}
