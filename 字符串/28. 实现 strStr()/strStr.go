package main

import (
    "fmt"
)


// 思路 1. 直接思路，暴力解法
func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    var ptr, j int
    for i := 0; i < len(haystack); i++ {
        // 第一个字母相同处开始匹配
        if haystack[i] == needle[j] {
            ptr = i
            // 第二个指针向后扫描
            for ptr < len(haystack) && j < len(needle) && haystack[ptr] == needle[j] {
                ptr++
                j++
            }
            // 根据匹配到的子串长度判断是否完成
            if ptr-i == len(needle) {
                return i
            } else {
                j = 0
            }
        }
    }
    return -1
}

// 思路 1优化. Sunday 算法
func strStr2(haystack string, needle string) int {}

func main() {
    haystack, needle := "aaaaa", "bba"
    result := strStr(haystack, needle)
    fmt.Println(result)
}