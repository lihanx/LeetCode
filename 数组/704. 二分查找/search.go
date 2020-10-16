package main

import (
	"fmt"
)

/*
704. 二分查找
https://leetcode-cn.com/problems/binary-search/
*/

func search(nums []int, target int) int {
	start, end := 0, len(nums) - 1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}

func main() {
	input := []int{5}
	res := search(input, 5)
	fmt.Println(res)
}