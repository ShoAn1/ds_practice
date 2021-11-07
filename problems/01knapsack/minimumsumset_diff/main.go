package main

import (
	"fmt"
	"math"
)

func subsetsum_minimum_diff() int {
	nums := [2]int{1, 4}
	sum := 0

	for _, i := range nums {
		sum += i
	}

	dp := make([][]bool, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i <= len(nums); i++ {
		for j := 0; j <= sum; j++ {
			if j == 0 {
				dp[i][j] = true
			} else if i == 0 {
				dp[i][j] = false
			} else {
				if nums[i-1] > j {
					dp[i][j] = dp[i-1][j]
				} else {
					if dp[i-1][j] || dp[i-1][j-nums[i-1]] {
						dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
					} else {
						dp[i][j] = dp[i-1][j]
					}
				}
			}
		}
	}
	fmt.Println(dp)
	min := math.MaxInt32
	temp := 0
	for i := 0; i <= sum/2; i++ {
		if dp[len(nums)][i] == true {
			temp = sum - 2*(i)
			if temp < min {
				min = temp
			}
		}
	}
	return min
}

func count_for_number_subsets_for_difference(difference int) {
	nums := [2]int{1, 4}
	sum := 0

	for _, i := range nums {
		sum += i
	}
	end_num := (sum + difference) / 2
	dp := make([][]int, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]int, sum+1)
	}

	for i := 0; i <= len(nums); i++ {
		for j := 0; j <= end_num; j++ {
			if j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else {
				if nums[i-1] > j {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
				}
			}
		}
	}
	fmt.Println(dp)
}
func main() {
	count_for_number_subsets_for_difference(1)

}
