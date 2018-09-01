package main

import "golang.org/x/tour/pic"

//import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var r [][]uint8

	for i := 0; i < dy; i++ {
		var t []uint8
		for j := 0; j < dx; j++ {
			t = append(t, uint8(i^j))
		}
		r = append(r, t)
	}
	return r
}

func main() {
	pic.Show(Pic)
}
