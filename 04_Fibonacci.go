//4.函数闭包练习，Fibonacci数列
//change:简化了算法，利用两个变量计算sum
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	sum, x := 0, 1
	return func() int {
		sum, x = x+sum, sum
		return sum
	}
}

func main() {
	f := fibonacci()
	fmt.Println(0)
	for i := 1; i < 10; i++ {
		fmt.Println(f())
	}
}
