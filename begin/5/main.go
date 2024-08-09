package main

import (
	"fmt"
	"math"
)

// Дана длина ребра куба a. Найти объем куба V = a3 и площадь его
// поверхности S = 6·a2 .

func volumeAndSurface(a float64) (volume, S float64)  {
	volume = math.Pow(a, 3)
	S = 6 * math.Pow(a, 2)
	return
}



func main() {
	volume, S := volumeAndSurface(3)
	formatted := fmt.Sprintf("Volume of cube: %2.f, Surface: %2.f from value of 3", volume, S)
	fmt.Println(formatted)
}