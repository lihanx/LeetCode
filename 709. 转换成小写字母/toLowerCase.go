package main

import (
    "fmt"
    "strings"
)

/*
709. 转换成小写字母
https://leetcode-cn.com/problems/to-lower-case/
*/

func toLowerCase(str string) string {
    length := len(str)
    if length == 0 {
        return str
    }
    var result strings.Builder
    for i := 0; i < length; i++ {
        if str[i] <= 90 && str[i] >= 65 {
            result.WriteByte(str[i] + 32)
        } else {
            result.WriteByte(str[i])
        }
    }
    return result.String()
}

func main() {
    s := "Hello World"
    res := toLowerCase(s)
    fmt.Printf("'%s'\n", res)
}