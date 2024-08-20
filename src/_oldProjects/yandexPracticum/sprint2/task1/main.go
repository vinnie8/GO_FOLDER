package main

import (
	"fmt"
	
)

func main(){

	hundredSlice := make([]int, 100) //создаем слайс на 100 элементов
	hundredSlice=sliceFill(hundredSlice)//заполняем слайс
	fmt.Println(hundredSlice, cap(hundredSlice))

	hundredSlice=sliceCut(hundredSlice)//удаляем элементы из середины слайса
	fmt.Println(hundredSlice, cap(hundredSlice))

	hundredSlice=sliceRotate(hundredSlice)//разворачиваем слайс
	fmt.Println(hundredSlice, cap(hundredSlice))
}


//fills slice with numbers from 1 to slice length
func sliceFill(slice []int)([]int){
	
	for i,_ :=range slice{
		slice[i]=i+1
	}
	return slice
}
//delete elements from mid the slice
func sliceCut(slice []int)([]int){
	var i int = 10

	 slice=append(slice[:i], slice[i+80:]...)
	 return slice
}
//rotates slice from top to bottom
func sliceRotate(slice []int)([]int){
	for i:= range slice[:len(slice)/2]{
		//меняем значения местами
		slice[i], slice[len(slice)-i-1]=slice[len(slice)-i-1], slice[i]	
	}
	 return slice
}