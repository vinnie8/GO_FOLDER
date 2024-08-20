package main

import (
	"fmt"
)
//задаем тип содержащий фигуры
type figures int

	const(
		square figures = iota // квадрат
		circle // круг
		triangle // равносторонний треугольник
) 

func main() {

	var myFigure = circle 
	
	ar, ok := area(myFigure)
	if !ok {
		fmt.Println("Ошибка")
		return
	}
	//вычисляем площадь фигуры
	myArea := ar(3)
	fmt.Println(myArea)
}

	func area(f figures) (func(float64) float64, bool) {
		//определяем фигуру из переданной переменной 
		switch f {
		case 0:
			//считаем площадь фигуры через замыкания
			return func(x float64) float64 { return x * x }, true
		case 1:
			return func(x float64) float64 { return 3.142 * x * x }, true
		case 2:
			return func(x float64) float64 { return 0.433 * x * x }, true
		default:
			return nil, false
		}
	}
	