package main

import (
    "fmt"
    "strings"
)

/*
剑指 Offer 05. 替换空格
https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof/
*/

func replaceSpace(s string) string {
    var res strings.Builder
    for _, c := range s {
        if c == ' ' {
            res.WriteString("%20")
        } else {
            res.WriteRune(c)            
        }
    }
    return res.String()
}

func main() {
    input := "Hello World! "
    res := replaceSpace(input)
    fmt.Println(res)
}