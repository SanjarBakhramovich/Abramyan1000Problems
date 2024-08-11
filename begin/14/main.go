package main

import (
	"fmt"
	"math"
)

// Дана длина L окружности. Найти ее радиус R и площадь S круга,
// ограниченного этой окружностью, учитывая, что L = 2·π·R, S = π·R2 .
// В качестве значения π использовать 3.14.

func findRadius(length float64) float64 {
	return length / (math.Pi * 2) 
}

func findArea(radius float64) float64 {
	return math.Pi * math.Pow(radius, 2)
}

func main() {
	L := 5

	R := findRadius(float64(L)) 
	S := findArea(R) 

	fmt.Printf("Радиус круга %.2f\n", R)
	fmt.Printf("Площадь круга %.2f\n", S)
}