package main

import (
    "fmt"
)


/*
415. 字符串相加
https://leetcode-cn.com/problems/add-strings/
*/

func addStrings(num1 string, num2 string) string {
    length1, length2 := len(num1), len(num2)
    result := make([]byte, max(length1, length2)+1)
    ptr1, ptr2, ptrr := length1-1, length2-1, len(result)-1
    var carry, tmp1, tmp2 byte
    for ptrr > 0 {
        if ptr1 < 0 {
            tmp1 = 0
        } else {
            tmp1 = num1[ptr1] - 48
            ptr1--
        }
        
        if ptr2 < 0 {
            tmp2 = 0
        } else {
            tmp2 = num2[ptr2] - 48
            ptr2--
        }
        result[ptrr] = (tmp1 + tmp2 + carry) % 10 + 48
        carry = (tmp1 + tmp2 + carry) / 10
        ptrr--
    }
    if carry == 0 {
        return string(result[1:])
    }
    result[0] = 49
    return string(result)
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}

func main() {
    result := addStrings("1", "99999")
    fmt.Printf("%s\n", result)
}