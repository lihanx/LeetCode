package main

import (
    "fmt"
)

/*
1544. 整理字符串
https://leetcode-cn.com/problems/make-the-string-great/
*/

func makeGood(s string) string {
    stack := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        stack = append(stack, s[i])
        n := len(stack)
        if n >= 2 && (stack[n-1]+32 == stack[n-2] || stack[n-1] - 32 == stack[n-2]) {
            stack = stack[:n-2]
        }
    }
    return string(stack)
}

func main() {
    // input := "leEeetcode"
    input := "abBAcC"
    result := makeGood(input)
    fmt.Println(result)
}