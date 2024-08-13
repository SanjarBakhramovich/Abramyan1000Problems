// Даны два целых числа A и B (A < B). Найти сумму всех целых чисел
// от A до B включительно.

package main

import "fmt"

func printSumInts(A, B int) int {
	totalSum := A

	for i := A; i <= B; i++{
		totalSum += i
	}
	return totalSum
}

func main() {
	var A, B int

	fmt.Print("Введите число A: ")
	fmt.Scan(&A)

	fmt.Print("Введите число B: ")
	fmt.Scan(&B)
	fmt.Println("Сумма всех чисел от", A, "до", B, "включительно", printSumInts(A,B))
}