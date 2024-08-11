package main

import (
	"fmt"
	"math"
)

// Даны катеты прямоугольного треугольника a и b. Найти его гипоте-
// нузу c и периметр P:
// √
// c = a2 + b2 ,
// P = a + b + c.



func findHypoten(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func findPerim(a, b, c float64) float64{
	return a + b + c
}

func main() {
	fmt.Println(findHypoten(3, 4))
	fmt.Println(findPerim(2,3,4))
}