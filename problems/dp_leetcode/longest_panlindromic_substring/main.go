package main

import (
	"fmt"
	"strings"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func inisialise_mat_topdown(m, n int) (t [][]int) {
	for i := 0; i <= m; i++ {
		temp := make([]int, n+1)
		t = append(t, temp)
	}
	return
}
func palindromic_substring(str string) string {
	s1 := strings.Split(str, "")
	max := 0
	i_new, j_new := 0, 0
	srev := strings.Split(reverse(str), "")
	length := len(s1)
	t := inisialise_mat_topdown(length, length)
	for i := 1; i <= length; i++ {
		for j := 1; j <= length; j++ {
			if s1[i-1] == srev[j-1] && i != j {
				t[i][j] = 1 + t[i-1][j-1]
				if t[i][j] > max {
					max = t[i][j]
					i_new = i - 1
					j_new = j - 1
				}
			} else {
				t[i][j] = 0
			}
		}
	}
	lar_palindromic_str := ""
	lar_palindromic_str_rev := ""
	for i := max; i >= 1; i-- {
		lar_palindromic_str += s1[i_new]
		lar_palindromic_str_rev += srev[j_new]
		i_new--
		j_new--
	}
	fmt.Println(lar_palindromic_str, lar_palindromic_str_rev)
	return lar_palindromic_str
}

func main() {
	s := "aacabdkacaa"
	value := palindromic_substring(s)
	fmt.Println(value)
}
