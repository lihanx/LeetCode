package main

import (
    "fmt"
    "strings"
)

func isFlipedString(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    s3 := s2 + s2
    if strings.Contains(s3, s1) {
        return true
    }
    return false
}

func main() {
    // s1, s2 := "waterbottle", "erbottlewat"
    s1, s2 := "aa", "aba"
    res := isFlipedString(s1, s2)
    fmt.Println(res)
}