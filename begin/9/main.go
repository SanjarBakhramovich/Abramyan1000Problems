package main

import (
	"fmt"
	"math"
)

// Даны два неотрицательных числа a и b. Найти их среднее геометри-
// ческое, то есть квадратный корень из их произведения: √a·b.

func findRootProd(a, b uint64) uint64 {
	return uint64(math.Sqrt(float64(a*b)))
}

func main() {
	fmt.Println(findRootProd(16, 16))
}