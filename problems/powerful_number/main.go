package main

func check_number(num, sum int) int {
	// if (num+sum)/10==0{
	// 	return num+sum
	// }
	if num == 0 {
		return 1
	}
	return 2 * check_number(num-1)
}

func main() {
}
