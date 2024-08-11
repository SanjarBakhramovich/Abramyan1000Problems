package main

import (
	"fmt"
	"math"
)

// Даны два ненулевых числа. Найти сумму, разность, произведение и
// частное их модулей.

func findSum(a, b uint64) float64{
	return math.Abs(float64(a)) + math.Abs(float64(b))
}

func findDiff(a, b uint64) float64{
	return math.Abs(float64(a)) - math.Abs(float64(b))
}

func findProduct(a, b uint64) float64{
	return math.Abs(float64(a)) * math.Abs(float64(b))
}
// 
func findQuotient(a, b uint64) float64{
	a2 := math.Abs(float64(a))
	b2 := math.Abs(float64(b))

	return a2 / b2
}

func main() {
	sum := findSum(uint64(2), uint64(2))
	fmt.Printf("Сумма двух модулей, 2 и 2 равна = %.0f\n", float64(sum))

	diff := findDiff(uint64(2), uint64(2))
	fmt.Printf("Разница между двух чисел модулей, 2 и 2 равна = %.0f\n", float64(diff))

	prod := findProduct(uint64(2), uint64(2))
	fmt.Printf("произведение двух чисел модулей, 2 и 2 равна = %.0f\n", float64(prod))

	quo := findQuotient(uint64(2), uint64(2))
	fmt.Printf("частное двух чисел модулей, 2 и 2 равна = %.0f\n", float64(quo))
}