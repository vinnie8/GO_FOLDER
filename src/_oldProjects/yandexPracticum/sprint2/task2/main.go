package main

import (
	"fmt"
	
)

func main(){
//creating map
	var price = map[string]int{
		"хлеб" : 50,
		"молоко" : 100,
		"масло" : 200,
		"колбаса" : 500,
		"соль" : 20,
		"огурцы" : 200,
		"сыр" : 600,
		"ветчина" : 700,
		"буженина" : 900,
		"помидоры" : 250,
		"рыба" : 300,
		"хамон" : 1500,
	}
	
	var order = []string{"хлеб", "буженина", "сыр", "огурцы"}

	showValuesOverThan(price,500)
	fmt.Println()
	calculateOrder(price,order)
}
//показывает товары которые стоят больше 500
func showValuesOverThan(priceMap map[string]int,setValue int){
	for key, value := range	priceMap{
		if value>=setValue{
			fmt.Print(key,":", value)
			fmt.Print("\t")
		}
	}
}
//считает сумму покупки, сравнивает товары в корзине и мапе и суммирует
func calculateOrder(priceMap map[string]int, order []string){
	
	var sum int = 0

	for _, v:=range order{
		sum+=priceMap[v]
	}
	
	fmt.Println("стоимость покупки составляет:", sum)
}