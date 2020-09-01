package main

import (
    "fmt"
)

/*
剑指 Offer 04. 二维数组中的查找
https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
*/

func findNumberIn2DArray(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }
    rows, columns := len(matrix), len(matrix[0])
    row, column := 0, columns - 1
    var num int
    for row < rows && column >= 0 {
        num = matrix[row][column]
        if num == target {
            return true
        } else if num > target {
            column--
        } else {
            row++
        }
    }
    return false
}