package main

import "fmt"

// Даны длины ребер a, b, c прямоугольного параллелепипеда. Найти его
// объем V = a·b·c и площадь поверхности S = 2·(a·b + b·c + a·c).

func findVolume(a,b,c float64) float64 {
	return a * b * c
}

func findSurf(a,b,c float64)  float64 {
	S := 2 * ((a*b) + (b * c) + (a * c))
	return S
}
func main() {
	fmt.Printf("Volume is equal to %.f\n", findVolume(2, 4, 6))
	// 
	fmt.Printf("Area of surface is equal to %.f from values of 2,4,6\n", findSurf(2,4,6))
}