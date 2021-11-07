package main

import "fmt"

func peak_height(arr []int) int {
	for i := 1; i < len(arr)-1; i++ {
		if arr[i]>arr[i+1] && arr[i]>arr[i-1]{
			return i
		}
	}
	
}

func main() {
	for i:=0;i=5;i++{
		fmt.Println(i)
	}
	fmt.Println("wednesday 11.50,14 july 2021")
	
}
