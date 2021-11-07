package main

import (
	"fmt"
	"strings"
)

// func longestValidPalinderome(a []string) int {
// 	i, max, count, left := 0, 0, 0, 0
// 	// stack := []int{}
// 	for i := 0; i < len(a); i++ {
// 		fmt.Print(a[i])
// 		if a[i] == "(" && i < len(a)-1 {
// 			left += 1
// 		} else if a[i] == ")" && left > 0 {
// 			count += 2
// 			left -= 1
// 		} else if left == 0 && a[i] == ")" {
// 			count = 0
// 		}
// 		fmt.Println(count)
// 		if count > max {
// 			max = count
// 		}
// 		i++
// 		fmt.Println(max)
// 	}
// 	return max - 2*left
// }

func recursecheck(index, left, count int, max *int, current []string) int {
	if index <= len(current)-1 {
		fmt.Println(current[index])
		if current[index] == ")" && left != 0 {
			return 1
		} else if current[index] == "(" {
			tmp := recursecheck(index+1, left+1, count, max, current)
			if tmp == 1 {
				count = 2 + 1
				if count > *max {
					*max = count
				}
				return count
			}
			return count
		} else {
			return recursecheck(index+1, left, 0, max, current)
		}
	}
	return 0
}

func longestValidParentheses(s string) int {
	a := strings.Split(s, "")
	stack := []string{}
	index := []int{}
	length := []int{}
	for i := 0; i < len(a); i++ {
		if a[i] == "(" {
			stack = append(stack, a[i])
			index = append(index, i)
		} else if a[i] == ")" {
			if len(stack) != 0 && stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
				cur := index[len(index)-1]
				length = append(length, i-cur)
				index = index[:len(index)-1]
			} else {
				stack = append(stack, ")")
				index = append(index, i)
				//cur := index[len(index)-1]
			}
		}
	}
	fmt.Println(index)
	fmt.Println(length)
	max := 0
	for i := 0; i <= len(length)-1; i++ {
		if length[i] > max {
			max = length[i]
		}
	}
	return max * 2
}

// func longestValidPalinderome(str string) int {
// 	a := strings.Split(str, "")
// 	// i, max, count := 0, 0, 0
// 	// stack := []string{}
// 	max := recursecheck(a)
// 	return max
// }

func main() {
	//str := "()(()"
	str := ")()())"
	//	s1 := strings.Split(str, "")
	//var max *int
	max := 0
	// s1 := strings.Split(str, "")
	fmt.Println(longestValidParentheses(str))
	//recursecheck(0, 0, 0, &max, s1)
	fmt.Println(max)
}
