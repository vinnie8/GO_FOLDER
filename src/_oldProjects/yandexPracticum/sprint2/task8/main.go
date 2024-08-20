// вычисление времени работы функции с помощью defer
package main

import (
	"fmt"
	"time"
)

func main() {
	VeryLongTimeFunction()
}

func metricTime(start time.Time) {
	// функция Now() возвращает текущее время, а функция Sub возвращает разницу между двумя временными метками
	fmt.Println("Time eliminated: ", time.Now().Sub(start))
}

func VeryLongTimeFunction() {
	defer metricTime(time.Now()) // передаём в функцию metricTime значение текущего времени и откладываем её вызов до возврата
	// Какие-то долгие вычисления
	var slice []int
	for i := 0; i < 100001; i++ {
		slice = append(slice, i)
	//	time.Sleep(2) explicitly added to slow down function
	}
	fmt.Println(cap(slice))//show capacity of slice

}