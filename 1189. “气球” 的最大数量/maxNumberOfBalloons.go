package main

import "fmt"

// 思路1 map 实现
func maxNumberOfBalloons(text string) int {
    record := make(map[byte]int)
    for i := 0; i < len(text); i++ {
        record[text[i]]++
    }
    charset := "balon"
    record[108] /= 2
    record[111] /= 2
    result := len(text)
    for j := 0; j < len(charset); j++ {
        val, _ := record[charset[j]]
        if val < result {
            result = val
        }
    }
    return result
}

// 思路2 数组实现
func maxNumberOfBalloons2(text string) int {
    record := make([]int, 26)
    for i := 0; i < len(text); i++ {
        record[text[i]-'a']++
    }
    record['l'-'a'] /= 2
    record['o'-'a'] /=2
    result := len(text)
    for _, val := range "balon" {
        if record[val-'a'] < result {
            if record[val-'a'] == 0 {
                return 0
            }
            result = record[val-'a']
        }
    }
    return result
}

func main() {
    input := "balloo"
    result := maxNumberOfBalloons2(input)
    fmt.Println(result)
}