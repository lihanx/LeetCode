package main

import (
    "fmt"
    "strings"
)

func sortString(s string) string {
    var bucket [26]int
    var result strings.Builder
    length := 0
    // 初始化 bucket
    for _, c := range s {
        bucket[c-'a']++
    }
    for length < len(s) {
        // 正向从小到大
        for i := 0; i < 26; i++ {
            if bucket[i] > 0 {
                result.WriteByte('a'+byte(i))
                bucket[i]--
                length++
            }
        }
        // 反向从大到小
        for i := 25; i >= 0; i-- {
            if bucket[i] > 0 {
                result.WriteByte('a'+byte(i))
                bucket[i]--
                length++
            }
        }
        
    }
    return result.String()
}

func main() {
    input := "leetcode"
    result := sortString(input)
    fmt.Println(result)
}