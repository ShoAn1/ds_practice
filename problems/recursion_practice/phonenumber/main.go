package main

import "fmt"

func phoneNumber_to_string(m1 map[int][]string, value []int, l1 int, s1 string) {

	if len(value) == 0 {
		fmt.Println(s1)
		return
	}
	for i, k := range value {
		v1 := m1[k]
		tempvalue := make([]int, 0)
		if i == 0 {
			tempvalue = append(tempvalue, value[i+1:]...)
		} else if i < len(value)-1 {
			tempvalue = append(tempvalue, value[:i]...)
			tempvalue = append(tempvalue, value[i+1:]...)
			//fmt.Println("if", tempvalue)
		} else {
			tempvalue = append(tempvalue, value[:i]...)
			//fmt.Println("else", tempvalue)
		}
		for _, j := range v1 {
			str := s1 + j
			phoneNumber_to_string(m1, tempvalue, l1, str)
		}
	}
}

func main() {
	value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m1 := map[int][]string{1: {"a", "b", "c"}, 2: {"d", "e", "f"}, 3: {"g", "h", "i"}, 4: {"j", "k", "l"},
		5: {"m", "n", "o"}, 6: {"p", "q", "r", "s"}, 7: {"t", "u", "v"}, 8: {"w", "x", "y"}, 9: {"z"}}
	// m1:=
	s1 := ""
	phoneNumber_to_string(m1, value, len(value), s1)
}
