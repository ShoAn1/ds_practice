package main

import (
	"fmt"
)

//array -> duplicate elements[] {2,14,9,0,8,7,4,9,11,7,1,9}

//2 14 9 0 8 7 4 11 7
//stack -9 9 7 7
// output -> 2,14,0,8,4,11,1,9,9,9,7,7

func array_check(value []int) {

	m1 := make(map[int]int)
	for i := 0; i <= len(value)-1; i++ {
		tempvalue := []int{}
		if valu, ok := m1[value[i]]; ok {
			if i == len(value)-1 {
				tempvalue = append(tempvalue, value[:i]...)
				fmt.Println("tempvalue", tempvalue)
			} else {

				tempvalue = append(tempvalue, value[:i]...)
				tempvalue = append(tempvalue, value[i+1:]...)
			}
			m1[value[i]] = valu + 1
			value = tempvalue
			i -= 1
		} else {
			m1[value[i]] = 1
		}

	}

	fmt.Println("initial array", value)
	length := len(value)
	for i := 0; i < length; i++ {
		tempvalue := []int{}
		valu, _ := m1[value[i]]
		if valu > 1 {
			key := value[i]
			tempvalue = append(tempvalue, value[:i]...)
			tempvalue = append(tempvalue, value[i+1:]...)
			for j := 0; j < valu; j++ {
				tempvalue = append(tempvalue, key)
			}
			value = tempvalue
			length--
		}

	}
	fmt.Println("map", m1)
	fmt.Println("array", value)

}

func main() {
	arr := []int{2, 14, 9, 0, 8, 7, 4, 9, 11, 7, 1, 17, 9, 7}
	fmt.Println("Hello, playground")
	array_check(arr)
}
