package main 

import (
    "fmt"
    "strings"
)

/*
1507. 转变日期格式
https://leetcode-cn.com/problems/reformat-date/
*/

func reformatDate(date string) string {
    month := map[string]string {
        "Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06", 
        "Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
    }
    array := strings.Split(date, " ")
    var res strings.Builder
    res.WriteString(array[2])
    res.WriteString("-")
    res.WriteString(month[array[1]])
    res.WriteString("-")
    res.WriteString(dtod(array[0]))
    return res.String()
}

func dtod(s string) string {
    d := s[:len(s)-2]
    if len(d) == 1 {
        return "0" + d
    }
    return d
}

func main() {
    date := "20th Oct 2052"
    result := reformatDate(date)
    fmt.Println(result)
}