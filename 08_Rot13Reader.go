//8.rot13Reader 练习
//简化了switch／case的逻辑
package main

import (
	"io"
	"os"
	//"fmt"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (ro *rot13Reader) Read(by []byte) (n int, err error) {
	n, err = ro.r.Read(by) //题目要求从io.Reader里读取数据
	//fmt.Print(by)
	for i, b := range by {
		switch {
		case 'A' <= b && b <= 'M' || 'a' <= b && b <= 'm':
			b = b + 13
		case 'M' < b && b <= 'Z' || 'm' < b && b <= 'z':
			b = b - 13

		}
		by[i] = b
	}

	return n, err

}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
