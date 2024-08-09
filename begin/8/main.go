package main

import "fmt"

// Даны два числа a и b. Найти их среднее арифметическое: (a + b)/2.

func averAri(a, b float64) float64 {
	return (a + b ) / 2
}

func main() {
	fmt.Println(averAri(3, 4))
}