//7.Reader 练习
//x循环改为range
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (m MyReader) Read(b []byte) (i int, e error) {
	for x := range b {
		b[x] = 'A'
	}
	return
}

func main() {
	reader.Validate(MyReader{})
}
