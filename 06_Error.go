//6.Error 练习
//简化了if／else结构；还原了Sqrt()的输入参数
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v.", float64(e))
}

var z float64 = 1.0

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		var last_z = z
		z = z - (z*z-x)/(2*z)
		if math.Abs(last_z-z) <= 0.00001 {
			return z, nil
		}
		return Sqrt(x)
	}

}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
