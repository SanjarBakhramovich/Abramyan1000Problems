package main

import "fmt"

// For1. Даны целые числа K и N (N > 0). Вывести N раз число K.

func printNum (N, K int) {
	for i:=0; i < N; i++ {
		fmt.Println(K)
	}
}

func main() {
	var K, N int

	fmt.Print("Введите число K: ")
	fmt.Scan(&K)
	// 
	fmt.Print("Введите число N: ")
	fmt.Scan(&N)

	printNum(N,K)
}