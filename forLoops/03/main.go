// For3. Даны два целых числа A и B (A < B).
// Вывести в порядке убывания
// все целые числа, расположенные между A и B (не включая числа A и B),
// а также количество N этих чисел.

package main

import "fmt"

func innrLoopPrint(a, b int) (counter int) {
	// 3 7 
	// Print 6 5 4
	// N = 3 

	for i := b-1; i > a; i--{
		fmt.Println(i)
		counter++
	} 
	return counter
}

func main() {
	var A, B int

	fmt.Print("Введите число A: ")
	fmt.Scan(&A)

	fmt.Print("Введите число B: ")
	fmt.Scan(&B)

	if A < B {
		N := innrLoopPrint(A, B)
		fmt.Printf("Количество чисел: %d\n", N)
	} else {
		fmt.Println("Число А должно быть меньше числа B.")
	}
}

// 3 7
// 6 5 4
