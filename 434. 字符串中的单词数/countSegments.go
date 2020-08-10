package main

import (
    "fmt"
)

func countSegments(s string) int {
    if s == "" {
        return 0
    }
    var start, end int
    n := len(s) - 1
    var space byte = 32
    // 去除两端空格
    for n >= 0 && s[n] == space {
        n--
    }    
    for start < n && s[start] == space {
        start++
    }
    if n < start {
        return 0
    }
    // 初始化 end 
    end = start - 1
    // 统计单词数
    var count int
    for end + 1 <= n {
        for end + 1 <= n && s[end+1] != space {
            end++
        }
        count++
        for end + 1 <= n && s[end+1] == space {
            end++
        }
        start = end + 1
    }
    return count
}

func main() {
    // input := "Hello, my name is John"
    // input := "          "
    // input := "a"
    // input := ", , , ,        a, eaefa"
    input := "    foo    bar"
    result := countSegments(input)
    fmt.Println(result)
}