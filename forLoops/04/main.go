// For4. Дано вещественное число — цена 1 кг конфет.
// Вывести стоимость 1,
// 2, . . . , 10 кг конфет.

package main

import "fmt"

func printPrice(price float64) float64{
	var totalCost float64
	for i := 1; i <= 10; i++{
		totalCost = price * float64(i)
		fmt.Printf("Цена за %d кг равно %.2f\n", i, totalCost)
	}
	return totalCost
}

func main() {
	var price float64
	fmt.Printf("Введите цену за 1 кг конфет: ")
	fmt.Scan(&price)

	printPrice(price)
}