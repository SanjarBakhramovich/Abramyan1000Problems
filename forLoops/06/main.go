// For6. Дано вещественное число — цена 1 кг конфет. Вывести стоимость 1.2,
// 1.4, . . . , 2 кг конфет.

package main

import "fmt"

func printFloatPrice(price float64) float64{
	var totalCost float64
	for i := 1.2; i <= 2; i += 0.2{
		totalCost = price * float64(i)
		fmt.Printf("Цена за %.1f кг конфет стоит %.2f единиц\n", i, totalCost)
	}
	return totalCost
}

func main() {
	var price float64
	fmt.Print("Введите цену за 1.2-кг конфет: ")
	fmt.Scan(&price)

	printFloatPrice(price)
}