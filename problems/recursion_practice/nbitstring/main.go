package main

import "fmt"

func calculate(n int, arr [4]int) {
	if n < 1 {
		fmt.Println(arr)
		return
	}
	arr[n-1] = 0
	calculate(n-1, arr)
	arr[n-1] = 1
	calculate(n-1, arr)
}

func main() {
	n := 4
	var arr [4]int
	calculate(n, arr)
}
