package main

import "math"

// Даны два круга с общим центром и радиусами R1 и R2 (R1 > R2 ).
// Найти площади этих кругов S 1 и S 2 ,
// а также площадь S 3 кольца,
// внешний радиус которого равен R1 ,
// а внутренний радиус равен R2 :
// S 1 = π·(R1 )2 ,
// S 2 = π·(R2 )2 ,
// S3 = S1 − S2.
// В качестве значения π использовать 3.14.

func findCircArea(radius float64) float64{
	return math.Pi * math.Pow(radius, 2)
}

func findOuterRadius()

func findInnerRadius()

func main() {
	
}