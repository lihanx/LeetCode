package main

import (
    "fmt"
)

/*
520. 检测大写字母
https://leetcode-cn.com/problems/detect-capital/
*/

func detectCapitalUse(word string) bool {
    var upper, lower int
    for _, v := range word {
        if isCapital(byte(v)) {
            upper++
        } else {
            lower++
        }
    }
    if upper == len(word) {
        return true
    }
    if lower == len(word) {
        return true
    }
    if upper == 1 && isCapital(word[0]) {
        return true
    }
    return false
}

func isCapital(b byte) bool {
    if b >= 'A' && b <= 'Z' {
        return true
    }
    return false
}

func main() {
    result := detectCapitalUse("USa")
    fmt.Println(result)
}