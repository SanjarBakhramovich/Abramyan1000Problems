// For2. Даны два целых числа A и B (A < B).
// Вывести в порядке возрастания все целые числа,
// расположенные между A и B (включая сами числа A и B),
// а также количество N этих чисел.

package main

import "fmt"

func printRangeWithCount(A, B int) (Counter int){
	for i := A; i <= B; i++ {
		fmt.Println(i)
		Counter++
	}
	return Counter
}

func main() {
	var A, B int

	fmt.Print("Введите число А: ")
	fmt.Scan(&A)
	// 
	fmt.Print("Введите число B: ")
	fmt.Scan(&B)	
	// 
	Counter := printRangeWithCount(A, B)
	fmt.Printf("Количество чисел в диапазоне: %d\n", Counter)
}