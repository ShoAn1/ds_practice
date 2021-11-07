package main

import (
	"fmt"
)

func pattern(n int) [5][5]int {
	a := [5][5]int{}
	front := 0
	back := 4
	n = 3
	for n != 0 {
		for i := front; i <= back; i++ {
			for j := front; j <= back; j++ {
				if i == front || i == back || j == front || j == back {
					a[i][j] = n
					//fmt.Println(a)
				}
			}
		}
		front += 1
		back -= 1
		n -= 1
	}
	fmt.Println(a)
	return a
}

func main() {
	//a := [5][5]int{}
	fmt.Println("Hello, playground")
	pattern(4)

}
