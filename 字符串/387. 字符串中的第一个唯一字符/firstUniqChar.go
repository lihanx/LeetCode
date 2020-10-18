package main

import (
	"fmt"
)

/*
387. 字符串中的第一个唯一字符
https://leetcode-cn.com/problems/first-unique-character-in-a-string/
*/

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	counter := make([]int, 26)
	for _, c := range s {
		counter[c-97] += 1
	}
	for i, c := range s {
		if counter[c-97] == 1 {
			return i
		}
	}
	return -1
}


func main() {
	input := "loveleetcode"
	res := firstUniqChar(input)
	fmt.Println(res)
}