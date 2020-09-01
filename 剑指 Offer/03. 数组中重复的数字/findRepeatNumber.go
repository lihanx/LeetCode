package main

import (
    "fmt"
)

/*
剑指 Offer 03. 数组中重复的数字
https://leetcode-cn.com/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
*/

func findRepeatNumber(nums []int) int {
    seen := make(map[int]struct{})
    for _, i := range nums {
        _, isSeen := seen[i]
        if isSeen {
            return i
        }
        seen[i] = struct{}{}
    }
    return -1
}


func main() {
    input := []int{2, 1, 4, 3, 1, 2}
    result := findRepeatNumber(input)
    fmt.Println(result)
}