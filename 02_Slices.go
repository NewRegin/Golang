//练习2:slices
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var image_size = make([][]uint8, dy)
	for i := range image_size {
		image_size[i] = make([]uint8, dx)
		for j := range image_size[i] {
			image_size[i][j] = uint8(dx+dy) / 2
		}
	}
	return image_size
}

func main() {
	pic.Show(Pic)
}
