package main

import "golang.org/x/tour/pic"
import "fmt"
import "math"

func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx, dy)
	var pict = make([][]uint8, dx)
	for i := 0; i < dy; i++ {
		pict[i] = make([]uint8, dy)
	}
	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			pict[x][y] = uint8(math.Sin(float64(x))+math.Cos(float64(y))/2)
		}
	}
	return pict
}

func main() {
	pic.Show(Pic)
}
