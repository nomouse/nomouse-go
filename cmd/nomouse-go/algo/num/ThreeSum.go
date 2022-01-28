package main

import (
	"fmt"
	"sort"
)

func threeNums(nums []int) []int {

	sort.Ints(nums)

	var size = len(nums)

	for i := 0; i < size-2; i++ {
		left, right := i+1, size-1

		for left < right {
			fmt.Println(i, "|", nums[left], "|", nums[right])
			diff := nums[left] + nums[right] + nums[i]

			if diff == 0 {
				return []int{nums[i], nums[left], nums[right]}
			} else if diff > 0 {
				right--
			} else {
				left++
			}
		}
	}

	//return nums
	return []int{}
}

func main() {

	nums := []int{-1, 4, -2, 7, 10, -3, 8, 11, 6}

	var result = threeNums(nums)
	fmt.Println(result)
}
