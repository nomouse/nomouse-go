package main

import "fmt"

func twoSum(nums []int, target int) []int {
	Map := make(map[int]int, 8)
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		diff := target - cur
		index, ok := Map[diff]
		fmt.Println(cur, "|", diff)
		if ok {
			return []int{index, i}
		} else {
			Map[cur] = i
		}
	}
	return []int{}
}

func main() {

	var nums = []int{-2, 2, 3, 4, 5, 10}
	var target int = 6

	fmt.Println(TwoSum(nums, target))
	//fmt.Println(target)
}
