package main

import (
    "fmt"
    "strings"
)

/*
剑指 Offer 58 - II. 左旋转字符串
https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
*/

func reverseLeftWords(s string, n int) string {
    var res strings.Builder
    for i := n; i < len(s); i++ {
        res.WriteByte(s[i])
    }
    for i := 0; i < n; i++ {
        res.WriteByte(s[i])
    }
    return res.String()
}


func main() {
    s, n := "abcdefg", 2
    result := reverseLeftWords(s, n)
    fmt.Println(result)
}