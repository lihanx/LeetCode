package main

import "fmt"

/*
58. 最后一个单词的长度
https://leetcode-cn.com/problems/length-of-last-word/
*/

func lengthOfLastWord(s string) int {
    n := len(s) - 1
    var space byte = 32
    i := n
    for i >= 0 {
        if s[i] == space {
            return n - i
        }
        i--
    }
    return 0
}

func main() {
    s := " "
    length := lengthOfLastWord(s)
    fmt.Println(s, length)
}