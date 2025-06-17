package main

import "fmt"

func twoSum(nums []int, target int) []int {
	nums_index := make(map[int]int)
	for i, v := range nums {
		complement, ok := nums_index[target-v]
		if ok {
			return []int{i, complement}
		}
		nums_index[v] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Println(result)
}
