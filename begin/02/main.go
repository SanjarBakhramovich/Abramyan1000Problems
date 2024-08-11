package main

import (
	"fmt"
	"math"
)

// Дана сторона квадрата a. Найти его площадь S = a2 .

func area(a float64) float64 {
	return math.Pow(a, 2)
}
func main() {
	fmt.Println(area(4))
}