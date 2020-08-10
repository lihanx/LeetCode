package main

import (
    "fmt"
)

func reverseVowels(s string) string {
    // 初始化双指针
    i, j := 0, len(s) - 1
    ba := []byte(s)
    for i < j {
        // ptr1 向后检测元音字母
        for i < len(s) && !isVowal(ba[i]) {
            i++
        }
        // ptr2 向前检测元音字母
        for j >= 0 && !isVowal(ba[j]) {
            j--
        }
        // 交换
        if i < j {
            ba[i], ba[j] = ba[j], ba[i]            
        }
        // 重新初始化避免原地死循环
        i++
        j--
    }
    return string(ba)
}

func isVowal(b byte) bool {
    switch b {
    case 97: return true
    case 101: return true
    case 105: return true
    case 111: return true
    case 117: return true
    case 65: return true
    case 69: return true
    case 73: return true
    case 79: return true
    case 85: return true
    default:
        return false
    }
}

func main() {
    // input := "leetcode"
    // input := "oe"
    // input := "hello"
    input := ".,"
    result := reverseVowels(input)
    fmt.Println(result)
}