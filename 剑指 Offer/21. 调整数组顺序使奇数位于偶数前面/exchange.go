package main

import (
	"fmt"
)

/*
剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
*/


/*
考虑双指针对向遍历，遍历过程中进行交换
那么有以下几种情况：
	1. nums[i], nums[j] 均为奇数
	2. nums[i], nums[j] 均为偶数
	3. nums[i] 为奇数, nums[j] 为偶数
	4. nums[i] 为偶数, nums[j] 为奇数

由于需要将奇数向前交换，那么
	当nums[i] 为奇数时，i++;
	当nums[j] 为偶数时，j--;
	当 i < j, 且 nums[i] 为偶数, nums[j] 为奇数时，进行交换;
*/

func exchange(nums []int) []int {
	i, j := 0, len(nums) - 1
	for i < j {
		if nums[i] % 2 == 0 && nums[j] % 2 != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		if nums[i] % 2 != 0 {
			i++
		}
		if nums[j] % 2 == 0 {
			j--
		}

	}
	return nums
}


func main() {
	nums := []int{11,9,3,7,16,4,2,0}
	result := exchange(nums)
	fmt.Println(result)
}