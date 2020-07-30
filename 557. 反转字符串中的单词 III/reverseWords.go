package main

import (
    "fmt"
    "strings"
)

/*
557. 反转字符串中的单词 III
https://leetcode-cn.com/problems/reverse-words-in-a-string-iii/
*/

func reverseWords(s string) string {  
    n := len(s) - 1
    
    if s == " " {
        return ""
    }

    var result strings.Builder
    var space byte = 32
    start, end := 0, -1
    
    for end < n {
        for end < n && s[end+1] != space {
            end++
        }
        for i := end; i >= start; i-- {
            result.WriteByte(s[i])
        }
        if end < n {
            result.WriteByte(space)
        }
        end++
        start = end + 1
    }
    return result.String()
}


func main() {
    s := "Let's take leetcode contest"
    // s := "a"
    rs := reverseWords(s)
    fmt.Println(rs)
}