package main

import "fmt"

// Дана сторона квадрата a. Найти его площадь S = a2 .

func area(a float32) float32 {
	return a * 2
}
func main() {
	fmt.Println(area(4))
}