package main

import "fmt"

func napsack(weight, value [4]int, knapsack, n int) [5][8]int {
	dp := [5][8]int{}
	for i := 0; i <= n; i++ {
		for j := 0; j <= 7; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				if weight[i-1] > j {
					dp[i][j] = dp[i-1][j]
				} else {
					if dp[i-1][j] < value[i-1]+dp[i-1][j-weight[i-1]] {
						dp[i][j] = value[i-1] + dp[i-1][j-weight[i-1]]
					} else {
						dp[i][j] = dp[i-1][j]
					}
				}
			}
		}
	}
	return dp
}
func main() {
	weight := [4]int{1, 3, 4, 5}
	value := [4]int{1, 4, 5, 7}
	val := napsack(weight, value, 17, 4)
	fmt.Println(val)
}
