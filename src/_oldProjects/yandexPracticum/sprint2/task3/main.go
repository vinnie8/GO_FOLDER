/*Дан массив целых чисел A и целое число k. 
Нужно найти и вывести индексы пары чисел, 
сумма которых равна k. Если таких чисел нет, 
то вернуть пустой слайс. 
Индексы можно вернуть в любом порядке.*/

package main

import (
	"fmt"
)

func main(){

var (
	A = map[int]int{
	1:10,
	2:20,
	3:14,
	4:56,
	5:25,
	6:6,
	7:118,
	8:1,
	9:3,
	10:69,
	11:50,
	12:45,
	}
	k int = 70
)

answer:=findPairOfIndex(A,k)

fmt.Println(answer)
}

func findPairOfIndex(A map[int]int, k int)(answer []int){

	var slice []int
	
	for i:=1;i<len(A);i++{
		for j:=i+1;j<len(A);j++{
			sum:=A[i]+A[j]
			
			if sum==k{
				slice = append(slice,i,j)
			}
		}
	}
return slice

}