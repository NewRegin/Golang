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
	for i := range by {
		switch {
		case 'A' <= by[i] && by[i] <= 'M' || 'a' <= by[i] && by[i] <= 'm':
			by[i] += 13
		case 'M' < by[i] && by[i] <= 'Z' || 'm' < by[i] && by[i] <= 'z':
			by[i] -= 13

		}
		
	}

	return n, err

}
func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
