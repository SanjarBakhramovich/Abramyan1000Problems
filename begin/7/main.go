package main

import (
	"fmt"
	"math"
)

// Найти длину окружности L и площадь круга S заданного радиуса R:
// L = 2·π·R,
// S = π·R2 .
// В качестве значения π использовать 3.14.

func findLength(R float64) float64{
	L := 2 * math.Pi * R
	return L
}

func findArea(R float64) float64 {
	S := math.Pi * math.Pow(R, 2)
	return S
}
func main() {
	fmt.Printf("Length %.2f\n", findLength(3))
	fmt.Printf("Area %.2f\n", findArea(3))
}