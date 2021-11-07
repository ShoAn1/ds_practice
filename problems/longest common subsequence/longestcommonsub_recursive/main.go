package main

import "fmt"

var t [][]int

func lcs_recursive(s1, s2 []rune, n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	if s1[n-1] == s2[m-1] {
		return 1 + lcs_recursive(s1, s2, n-1, m-1)
	} else {
		v1, v2 := lcs_recursive(s1, s2, n-1, m), lcs_recursive(s1, s2, n, m-1)
		if v1 > v2 {
			return v1
		}
		return v2
	}
}

func lcs_recur_meoization(s1, s2 []rune, n, m int) int {
	inisialise_mat_memo(n, m)
	if n == 0 || m == 0 {
		return 0
	}
	if t[n][m] != -1 {
		return t[n][m]
	}
	if s1[n-1] == s2[m-1] {
		t[n][m] = 1 + lcs_recursive(s1, s2, n-1, m-1)
		return t[n][m]
	} else {
		v1, v2 := lcs_recursive(s1, s2, n-1, m), lcs_recursive(s1, s2, n, m-1)
		if v1 > v2 {
			t[n][m] = v1
			return v1
		}
		t[n][m] = v2
		return v2
	}
}

func lcs_topDown() {

}
func inisialise_mat_topdown(m, n int) {
	for i := 0; i <= m; i++ {
		temp := make([]int, n)
		t = append(t, temp)
	}
}

func inisialise_mat_memo(m, n int) {
	for i := 0; i <= m; i++ {
		temp := make([]int, 0)
		for i := 0; i <= n; i++ {
			temp = append(temp, -1)
		}
		t = append(t, temp)
	}
}

func main() {
	s1 := []rune("abcdaf")
	s2 := []rune("acbcf")
	//fmt.Println(lcs(s1, s2, len(s1), len(s2)))
	// inisialise_mat(len(s1), len(s2))
	for _, i := range t {
		fmt.Println(i)
	}
	lcs_recur_meoization(s1, s2, len(s1), len(s2))
	fmt.Println("lcs value ", t[len(s1)][len(s2)])
}
