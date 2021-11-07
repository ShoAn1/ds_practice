package main

import "fmt"

func nlengthnumber(n int, arr [6]int) {
	if n < 1 {
		fmt.Println(arr)
	} else {
		for i := 1; i <= 6; i++ {
			arr[n-1] = i
			nlengthstring(n-1, arr)
		}
	}
}

func nlength_string(arr []string, i, j int) {
	if n == i {
		return
	} else {
		for _, i := range arr {
			arr := append(arr, i)
		}
	}
}

func main() {
	var a [6]int
	arr := []string("shouryaanand")
	fmt.Println(nlength_string(6, arr))

}
