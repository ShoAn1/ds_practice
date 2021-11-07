package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// func compareTime(s1, s2 string) {
// 	hr1 := strings.Split(s1, ":")
// 	hr2 := strings.Split(s2, ":")

// }

// func sortTime(arr []string) int {
// 	for i := 0; i <= len(arr)-1; i++ {
// 		for j := i + 1; j <= len(arr)-1; j++ {
// 			hr2 := strings.Split(arr[j], ":")
// 			h2, _ := strconv.Atoi(hr2[0])
// 			m2, _ := strconv.Atoi(hr2[1])
// 			hr1 := strings.Split(arr[i], ":")
// 			h1, _ := strconv.Atoi(hr1[0])
// 			m1, _ := strconv.Atoi(hr1[1])
// 			if h1 < h2 || (h1 == h2 && m1 <= m2) {
// 				arr[i], arr[j] = arr[j], arr[i]
// 				fmt.Println(arr[i], arr[j])
// 			}
// 		}
// 	}
// 	return arr
// }

func convertminute(s1 string) int {
	hr2 := strings.Split(s1, ":")
	h2, _ := strconv.Atoi(hr2[0])
	m2, _ := strconv.Atoi(hr2[1])
	return (h2*60 + m2)
}
func getMaxAndMin(time []string) string {
	largest := convertminute(time[0])
	smallest := convertminute(time[0])
	for _, i := range time {
		if convertminute(i) > (largest) {
			largest = convertminute(i)
		}
		if convertminute(i) < (smallest) {
			smallest = convertminute(i)
		}
	}
	fin_hour := strconv.Itoa((largest - smallest) / 60)
	fin_minutes := strconv.Itoa((largest - smallest) % 60)
	return fin_hour + ":" + fin_minutes
}

func GetMinimumNew(s1 []string) int {
	min := 24 * 60
	arr := []int{}
	for i := 0; i <= len(s1)-1; i++ {
		arr = append(arr, convertminute(s1[i]))
	}
	sort.Ints(arr)
	fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		if min > arr[i]-arr[i-1] {
			min = arr[i] - arr[i-1]
		}
		if min > 1440-(arr[i]-arr[i-1]) {
			min = 1440 - (arr[i] - arr[i-1])
		}
	}
	if min > arr[len(arr)-1]-arr[0] {
		min = arr[len(arr)-1] - arr[0]
	}
	if min >= 1440-(arr[len(arr)-1]-arr[0]) {
		min = 1440 - (arr[len(arr)-1] - arr[0])
	}
	return min

}

// func getMinimumminutes(s1 []string) int {
// 	smallest := int64(24 * 60)
// 	for i := 0; i <= len(s1)-1; i++ {
// 		for j := i + 1; j <= len(s1)-1; j++ {
// 			if smallest > int64(math.Abs(float64(convertminute(s1[i])-convertminute(s1[j])))) {
// 				smallest = int64(math.Abs(float64(convertminute(s1[i]) - convertminute(s1[j]))))
// 			}
// 			if int64(smallest) > int64(24*60-math.Abs(float64(convertminute(s1[i])-convertminute(s1[j])))) {
// 				smallest = int64(24*60 - math.Abs(float64(convertminute(s1[i])-convertminute(s1[j]))))
// 			}
// 		}
// 	}
// 	return int(smallest)
// }
func main() {
	// arr := []int{5, 4, 3, 1, 2, 7}
	time := []string{"12:12", "12:13", "00:12", "00:13"}
	sort_time := GetMinimumNew(time)

	fmt.Println(sort_time)

}
