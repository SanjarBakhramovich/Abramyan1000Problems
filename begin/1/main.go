package main

import "fmt"

// Дана сторона квадрата a. Найти его периметр P = 4·a.
func Perimeter(storona float64) float64 {
	return storona * 4
}

func main() {
	fmt.Println(Perimeter(4))
}