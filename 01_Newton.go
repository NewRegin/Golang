//1.1作为练习函数和循环的简单途径，用牛顿法实现开方函数。
//修改了递归和返回的顺序；迭代法可以考虑for循环知道满足条件math.Abs(c-z) <= 0.00001即可。
package main

import (
	"fmt"
	"math"
)

func Sqrt(x, z float64) float64 {
	var last_z = z
	z = z - (z*z-x)/(2*z)

	if math.Abs(last_z-z) <= 0.00001 {
		return z
	}
	return Sqrt(x, z)
}

func main() {
	fmt.Println(Sqrt(2, 1))
}
