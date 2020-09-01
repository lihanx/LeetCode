package main

import (
    "fmt"
)

/*
1436. 旅行终点站
https://leetcode-cn.com/problems/destination-city/
*/


func destCity(paths [][]string) string {
    record := make(map[string]int)
    destCitys := make([]string, 0)
    for _, row := range paths {
        for _, city := range row {
            record[city]++
        }
        destCitys = append(destCitys, row[len(row)-1])
    }
    var destCity string
    for _, city := range destCitys {
        val, _ := record[city]
        if val == 1 {
            destCity = city
        }
    }
    return destCity
}

func main() {
    input := [][]string{
        []string{"jMgaf WaWA","iinynVdmBz"},
        []string{" QCrEFBcAw","wRPRHznLWS"},
        []string{"iinynVdmBz","OoLjlLFzjz"},
        []string{"OoLjlLFzjz"," QCrEFBcAw"},
        []string{"IhxjNbDeXk","jMgaf WaWA"},
        []string{"jmuAYy vgz","IhxjNbDeXk"},
    }
    result := destCity(input)
    fmt.Println(result)
}