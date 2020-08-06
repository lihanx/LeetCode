package main

import "fmt"

/*
917. 仅仅反转字母
https://leetcode-cn.com/problems/reverse-only-letters/
*/

// 思路 1
func reverseOnlyLetters(S string) string {
    stack := make([]byte, 0)
    result := make([]byte, 0)
    // 初始化栈
    for i := 0; i < len(S); i++ {
        if isAlpha(S[i]) {
            stack = append(stack, S[i])
        }        
    }
    // 遍历反转
    ptr := len(stack) - 1
    for j := 0; j < len(S); j++ {
        if isAlpha(S[j]) {
            result = append(result, stack[ptr])
            ptr--
        } else {
            result = append(result, S[j])
        }
    }
    return string(result)
}

// 思路 2
func reverseOnlyLetters2(S string) string {
    result := make([]byte, 0)
    ptr := len(S) - 1
    for i := 0; i < len(S); i++ {
        if isAlpha(S[i]) {
            for !isAlpha(S[ptr]) {
                ptr--
            }
            result = append(result, S[ptr])
            ptr--
        } else {
            result = append(result, S[i])
        }
    }
    return string(result)
}

// 根据 asciiCode 判断是否为字母
func isAlpha(b byte) bool {
    if b >= 65 && b <= 90 {
        return true
    }
    if b >= 97 && b <= 122 {
        return true
    }
    return false
}

func main() {
    input := "Test1ng-Leet=code-Q!"
    result1 := reverseOnlyLetters(input)
    result2 := reverseOnlyLetters2(input)
    fmt.Printf("%s\n%s\n%s\n", input, result1, result2)
}