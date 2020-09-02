package main

/*
剑指 Offer 10- I. 斐波那契数列
https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof/
*/

func fib(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, (a+b)%(1e9+7)
	}
	return a 
}


func main()