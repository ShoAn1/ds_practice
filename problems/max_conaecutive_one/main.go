package main

import "fmt"

func longestOnes(nums []int, k int) int {
	left := 0
	right := 0

	for ; right < len(nums); right++ {
		if nums[right] == 0 {
			k--
		}
		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}
	}

	return right - left
}

func main() {
	fmt.Println("Date is 1 July 2021")
	arr := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k := 3
	value := longestOnes(arr, k)
	fmt.Println("returned value ", value)
}
