package main

import "fmt"

// Даны стороны прямоугольника a и b. Найти его площадь S = a·b и
// периметр P = 2·(a + b).

// find area and perimeter of square

func area(a, b float32) float32 {
	return a * b
}

func perim(a, b float32) float32{
	return 2 *  (a + b)
}


func main() {
	fmt.Println(area(3, 4))
	fmt.Println(perim(5,6))
}