package main

import (
	"errors"
	"fmt"
)

func build_order(status map[int]int, deps_map map[int][]int, j int) (int, error) {
	deps, ok := deps_map[j]
	if ok {
		if status[j] == 2 || len(deps) == 0 {
			return 2, nil
		}
		if status[j] == 1 {
			return 0, errors.New("cyclic dependecy")
		}
		status[j] = 1
		var err error
		for _, i := range deps {
			status[i], err = build_order(status, deps_map, i)
			if err != nil {
				return 0, err
			}
		}
	}
	status[j] = 2
	deps_map[j] = []int{}
	fmt.Println(j)
	return 0, nil
}
func main() {
	val_dep := make(map[int][]int)
	status := make(map[int]int)
	input := []int{1, 2, 3, 4, 5}
	input_deps := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}}
	for _, j := range input_deps {
		if val, ok := val_dep[j[0]]; ok {
			val_dep[j[0]] = append(val, j[1])
		} else {
			val_dep[j[0]] = make([]int, 0)
			val_dep[j[0]] = append(val_dep[j[0]], j[1])
		}
	}
	for _, j := range input {
		_, err := build_order(status, val_dep, j)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
