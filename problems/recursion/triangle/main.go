package main

import "fmt"

func check_triangle(arr []int, sum, count int) int {
	check := 0
	if count == 2 {
		check := 0
		for _, i := range arr {
			if sum > i {
				check++
			}
		}
		return check
	} else {
		for i := 0; i < len(arr); i++ {
			sum = sum + arr[i]
			count++
			tmp := []int{}
			tmp = append(tmp, arr[:i]...)
			tmp = append(tmp, arr[i+1:]...)
			check += check_triangle(tmp, sum, count)
		}
	}
	return check
}

func main() {
	arr := []int{1, 5, 6, 7, 3, 9}
	fmt.Println(check_triangle(arr, 0, 0))

}
