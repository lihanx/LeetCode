package main

import (
    "fmt"
    "strings"
)

func toGoatLatin(S string) string {
    var space byte = 32
    var start, end int
    start, end = 0, -1
    count := 1
    var result strings.Builder
    n := len(S)
    // 去除尾部空格
    for S[n-1] == space {
        n--
    }
    for end + 1 < n {
        // 找到单词尾部
        for end + 1 < n && S[end+1] != space {
            end++
        }
        // 写入尾部字母
        if S[start] == 97 || S[start] == 101 || S[start] == 105 || S[start] == 111 || S[start] == 117 || S[start] == 65 || S[start] == 69 || S[start] == 73 || S[start] == 79 || S[start] == 85 {
            // 从头顺序写入字母
            for i := start; i <= end; i++ {
                result.WriteByte(S[i])
            }
        } else {
            // 将首字母移到最后
            for i := start+1; i <= end; i++ {
                result.WriteByte(S[i])
            }
            result.WriteByte(S[start])
        }
        // 添加Goat Latin后缀
        result.WriteString("ma")
        for i := 0; i < count; i++ {
            result.WriteString("a")
        }
        
        // 进到下一个单词的起始位置
        end++
        // 增加位置判断，避免在结尾写入空格
        if end + 1 < n {
            result.WriteByte(space)
        }
        start = end + 1
        // 记录单词序数
        count++
    }
    return result.String()
}

func main() {
    input := "Each word consists of lowercase and uppercase letters only"
    result := toGoatLatin(input)
    fmt.Printf("'%s'\n", result)
}