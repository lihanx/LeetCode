package main

import (
    "fmt"
    "strings"
)

func reverseWords(s string) string {
    if s == "" {
        return ""
    }
    var result strings.Builder
    var space byte = 32
    // 清除两端空格
    left, right := 0, len(s) - 1
    // 注意特殊情况，全空格或空字符，需要限制 index 范围
    for left < len(s) && s[left] == space {
        left++
    }
    for right >= 0 && s[right] == space {
        right--
    }
    start, end := right, right
    for start > left - 1 {
        // 寻找单词
        for start > left - 1 && s[start] != space {
            start--
        }
        // 写入单词
        result.WriteString(s[start+1:end+1])
        // 写入空格
        if start > left {
            result.WriteByte(space)
        }
        // 忽略中间多个字符
        for start > left - 1 && s[start] == space {
            start--
        }
        // 重置滑窗重点
        end = start
    }
    return result.String()
}


func main() {
    // input := "  hello world!  "
    // input := "a good   example"
    // input := "   "
    // input := ""
    input := "a"
    result := reverseWords(input)
    fmt.Printf("'%s'", result)
}