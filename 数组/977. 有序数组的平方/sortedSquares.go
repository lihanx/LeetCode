package main

import (
	"fmt"
)

/*
977. 有序数组的平方
https://leetcode-cn.com/problems/squares-of-a-sorted-array/
*/

func sortedSquares(A []int) []int {
	i, j := 0, len(A) - 1
	res := make([]int, len(A))
	p := len(A) - 1
	
	for i <= j {
		if abs(A[i]) > abs(A[j]) {
			res[p] = abs(A[i] * A[i])
			i++
		} else {
			res[p] = abs(A[j] * A[j])
			j--
		}
		p--
	}
	return res
}


func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}


func main() {
	input := []int{-7,-3,2,3,11}
	res := sortedSquares(input)
	fmt.Println(res)
}