/*Используя defer,
напишите функцию, которая меняет переменную Global и выводит 
на экран её новое значение. Потом она должна вернуть всё как было.*/

package main

import "fmt"

var Global = 5

func main(){

	changeValue(&Global)
	
}
func changeValue(global *int){
	
	defer func(){
		*global=5
		fmt.Println(Global)
	}()

	*global=6 //меняется переменная global
	fmt.Println(Global)
	//тут срабатывает отложенная функция и возвращает значение обратно
}

