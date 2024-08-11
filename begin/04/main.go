package main

import (
	"fmt"
	"math"
)

// Дан диаметр окружности d. Найти ее длину L = π·d. В качестве
// значения π использовать 3.14.

func length(d float32) float32 {
	return d * math.Pi
}


func main() {
	fmt.Println(length(5))
}