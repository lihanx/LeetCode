package main

import "fmt"

/*
521. 最长特殊序列 Ⅰ
https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/

题解：
字符串 aa 和 bb 共有 3 种情况：

a=b。如果两个字符串相同，则没有特殊子序列，返回 -1。

length(a) == length(b) 且 a != b。例如：abcabc 和 abdabd。这种情况下，一个字符串一定不会是另外一个字符串的子序列，因此可以将任意一个字符串看作是特殊子序列，返回 length(a)length(a) 或 length(b)length(b)。

length(a) != length(b)。例如：abcdabcd 和 abcabc。这种情况下，长的字符串一定不会是短字符串的子序列，因此可以将长字符串看作是特殊子序列，返回 max(length(a),length(b))。

作者：LeetCode
链接：https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/solution/zui-chang-te-shu-xu-lie-i-by-leetcode/
来源：力扣（LeetCode）
*/

func findLUSlength(a string, b string) int {
    la, lb := len(a), len(b)
    if la == lb && a != b {
        return la
    }
    if la != lb {
        if la > lb {
            return la
        } else {
            return lb
        }
    }
    return -1
}

func main() {
    result := findLUSlength("aa", "aaa")
    fmt.Println(result)
}