package main

import (
    "fmt"
    "strconv"
)


func compress(chars []byte) int {
    var anchor, write int
    for read, c := range chars {
        if read + 1 == len(chars) || chars[read + 1] != c {
            char[write] = char[anchor]
            write++
            if read > anchor {
                for _, c := range strconv.itoa(read - anchor + 1) {
                    char[write] = c
                    write++
                }
            }
            anchor = read + 1
        }
    }
    return write
}