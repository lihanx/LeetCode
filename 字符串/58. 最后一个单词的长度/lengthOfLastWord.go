package main

import "fmt"

/*
58. 最后一个单词的长度
https://leetcode-cn.com/problems/length-of-last-word/
*/

func lengthOfLastWord(s string) int {
    
    length := len(s)
    if length == 0 {
        return 0
    }
    end := length - 1
    var space byte = 32
    // 过滤尾部空白字符
    for end >= 0 && s[end] == space {
        end--
    }
    // 滑窗起点
    start := end
    for start >= 0 && s[start] != space {
        start--
    }
    return end - start
}

func main() {
    s := "H"
    length := lengthOfLastWord(s)
    fmt.Printf("'%s', %d\n", s, length)
}