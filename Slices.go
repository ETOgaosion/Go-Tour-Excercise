package main

import "golang.org/x/tour/pic"
//import "math"

func pixelGen(x, y int) uint8 {
	return uint8(x * (y + 1))
}

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, 0, dy)
	for i := 0; i < dy; i++ {
		resi := make([]uint8, 0, dx)
		for j := 0; j < dx; j++ {
			resi = append(resi, pixelGen(i, j))
		}
		res = append(res, resi)
	}
	return res
}

func main() {
	pic.Show(Pic)
}
