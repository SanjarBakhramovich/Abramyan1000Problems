package main

import (
	"fmt"
	"math"
)

// Даны два ненулевых числа. Найти сумму, разность, произведение и
// частное их квадратов.

func findSum(a, b uint64) float64{
	return math.Pow(float64(a), 2) + math.Pow(float64(b), 2)
}

func findDiff(a, b uint64) float64{
	return math.Pow(float64(a), 2) - math.Pow(float64(b), 2)
}

func findProduct(a, b uint64) float64{
	return math.Pow(float64(a), 2) * math.Pow(float64(b), 2)
}
// 
func findQuotient(a, b uint64) float64{
	a2 := math.Pow(float64(a), 2)
	b2 := math.Pow(float64(b), 2)

	return a2 / b2
}

func main() {
	sum := findSum(uint64(2), uint64(2))
	fmt.Printf("Сумма квадратов двух чисел, 2 и 2 равна = %.0f\n", float64(sum))

	diff := findDiff(uint64(2), uint64(2))
	fmt.Printf("Разница между двух чисел, 2 и 2 равна = %.0f\n", float64(diff))

	prod := findProduct(uint64(2), uint64(2))
	fmt.Printf("произведение двух чисел, 2 и 2 равна = %.0f\n", float64(prod))

	quo := findQuotient(uint64(2), uint64(2))
	fmt.Printf("частное двух чисел, 2 и 2 равна = %.0f\n", float64(quo))
}