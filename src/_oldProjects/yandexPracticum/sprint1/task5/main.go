package main

import 	"fmt"

func main(){

var (
	a int=5 //integer
	p *int) //pointer

p=&a //adress getting

	fmt.Println(a,p)// смотрим на адресс в памяти на который ссылается указатель
}