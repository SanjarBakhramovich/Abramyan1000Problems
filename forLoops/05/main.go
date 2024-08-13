// For5*. Дано вещественное число — цена 1 кг конфет. Вывести стоимость 0.1,
// 0.2, . . . , 1 кг конфет.

package main

import "fmt"

func printFloatPrice(price float64) float64{
	var totalCost float64
	for i := 0.1; i <= 1; i+= 0.1{
		totalCost = price * float64(i)
		fmt.Printf("Цена за %.1f кг равно %.2f\n", i, totalCost)
	}
	return totalCost
}

func main() {
	var price float64
	fmt.Printf("Введите цену за 1-кг конфет: ")
	fmt.Scan(&price)

	printFloatPrice(price)

}