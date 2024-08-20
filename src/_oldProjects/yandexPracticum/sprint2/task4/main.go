//Напишите функцию, которая убирает дубликаты, 
//сохраняя порядок слайса
package main

import "fmt"

func main(){

	var input  = []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	fmt.Println(input)
	fmt.Println(RemoveDuplicates(input))

}

func RemoveDuplicates(inputSlice []string) []string {
//создаем пустой словарь
	m:=make(map[string]int)

	outputSlice:=make([]string,len(inputSlice))
//проверяем наличие в словаре элементов из входящего среза
	var index int=0
	for _, val:=range inputSlice{
		if _,ok:=m[val];!ok{
			m[val]=index //если элемента еще нет в словаре, 
			outputSlice[index]=val
			index++	 //заносим его в словарь и присваиваем номер по порядку
		} 
	}
	
return outputSlice[:index]
}
