package main

import (
    "bufio"
    "fmt"
    "os"
)
/*
Перед вами неполный код программы, которая считает, 
сколько строк пользователь ввёл в консоль, 
и после ввода каждой новой строки выводит 
общее количество на экран. Напишите реализацию функции f.*/

func main(){
	// Получаем читателя пользовательского ввода
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Interaction counter")

	cnt:=0
	
	for {
		fmt.Print("->")
		// Считываем введённую пользователем строку. 
		//Программа ждёт, пока пользователь введёт строку
		a, err:=reader.ReadString('\n')
		if err !=nil{
			panic(err)
		}

		f(&cnt)//реализовать счетчик строк

		fmt.Printf("User input %d lines\n",cnt)
	}
}

//реализация метода f с помощью указателя
func f(cnt *int) {
	*cnt++
}