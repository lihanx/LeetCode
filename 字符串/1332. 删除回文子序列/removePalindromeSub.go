package main

import (
    "fmt"
)

/*
1332. 删除回文子序列
https://leetcode-cn.com/problems/remove-palindromic-subsequences/
*/

func removePalindromeSub(s string) int {
    if s == "" {
        return 0
    }
    if isPalindrome(s) {
        return 1
    }
    return 2
}

func isPalindrome(s string) bool {
    i, j := 0, len(s) - 1
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func main() {
    // s := "ababa"
    s := "abba"
    res := removePalindromeSub(s)
    fmt.Println(res)
}