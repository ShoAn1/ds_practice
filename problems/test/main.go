package main

import "fmt"

func findElement(arr []int, start, length, element int) int {
	var left, right int
	if start < length {
		mid := (start + length) / 2
		//fmt.Println(mid)
		if arr[mid] == element {
			return mid
		}
		// if arr[mid] > element {
		// return findElement(arr, start, mid-1, element)
		// } else {
		left = findElement(arr, start, mid-1, element)
		right = findElement(arr, mid+1, length, element)
		// return findElement(arr, mid+1, length, element)
		// }
		if left > right {
			return right
		}
		return left
	}
	return -1
}

// func PrintArrayElements(array map[string]interface{}) []int {

// 	return arr
// }
func main() {
	arr := []int{5, 6, 7, 8, 9, 0, 1, 2, 3, 4}
	fmt.Println(findElement(arr, 0, len(arr)-1, 4))
}
