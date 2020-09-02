package main

import "fmt"

/*
剑指 Offer 11. 旋转数组的最小数字
https://leetcode-cn.com/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/

把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。  

示例 1：
	输入：[3,4,5,1,2]
	输出：1
示例 2：
	输入：[2,2,2,0,1]
	输出：0
*/

func minArray(numbers []int) int {
	start, end := 0, len(numbers) - 1
	for start < end {
		mid := (start + end) / 2
		if numbers[mid] > numbers[end] {
			start = mid + 1
		} else if numbers[mid] < numbers[end] {
			end = mid
		} else {
			end--
		}
	}
	return numbers[start]
}


func main() {
	// numbers := []int{3, 4, 5, 6, 1, 2, 3}
	numbers := []int{0}
	res := minArray(numbers)
	fmt.Println(res)
}