package main

import "fmt"
//Добавьте в пример формулу с iota,
// чтобы программа выводила 1 3 5 7 9 11
const(
	one = 1+iota*2
	three
	five
	seven
	nine
	eleven
)

func main(){

	fmt.Println(one,three,five,nine,eleven)

}
