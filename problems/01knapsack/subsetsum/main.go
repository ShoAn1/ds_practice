package main

import (
	"fmt"
)

// func napsack(weight [5]int, knapsack, n int) [6][12]bool {
// 	dp := [6][12]bool{}

// 	for i := 0; i <= n; i++ {
// 		for j := 0; j <= 11; j++ {
// 			if j == 0 {
// 				dp[i][j] = true
// 			} else if i == 0 {
// 				dp[i][j] = false
// 			} else {
// 				if weight[i-1] > j {
// 					dp[i][j] = dp[i-1][j]
// 				} else {
// 					if dp[i-1][j] || dp[i-1][j-weight[i-1]] {
// 						dp[i][j] = dp[i-1][j] || dp[i-1][j-weight[i-1]]
// 					} else {
// 						dp[i][j] = dp[i-1][j]
// 					}
// 				}

// 			}
// 		}
// 	}
// 	return dp
// }

func canPartition(nums []int) [][]bool {
	var sum int = 0
	for _, i := range nums {
		sum = sum + i
	}
	if sum%2 != 0 {
		return [][]bool{}
	}
	sum = sum / 2
	var length int = len(nums)
	var dp = make([][]bool, length+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i <= length; i++ {
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
				// if dp[i][j] {
				// 	return true
				// }
			}
		}
	}
	fmt.Println(dp[length][sum])
	return dp

}

func canPartitionKSubsets(nums []int, k int) [][]bool {
	var sum int = 0
	for _, i := range nums {
		sum = sum + i
	}
	if sum%k != 0 {
		return [][]bool{}
	}
	sum = sum / k
	var length int = len(nums)
	var dp = make([][]bool, length+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i <= length; i++ {
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
				// if dp[i][j] {
				// 	return true
				// }
			}
		}
	}
	count := 0
	for i, _ := range dp {
		if dp[i][sum] {
			count++
		}
	}
	fmt.Println(count)
	return dp
}

func main() {
	//	weight := [5]int{2, 3, 7, 8, 10}
	vvv := []int{1, 2, 2, 2, 2}
	// value := [4]int{1, 4, 5, 7}
	// val := napsack(weight, 11, 5)
	v1 := canPartitionKSubsets(vvv[:], 3)
	//fmt.Println("***************", v1)
	for _, i := range v1 {
		fmt.Println(i)
	}
}
