/*Напишите код,
 который будет сериализовывать структуру в json-строку следующего вида:*/
package main

import (
	"fmt"
	"encoding/json"
	"log"
	"time"
)

func main(){
//создаем экземпляр структуры
p:=Person{
	Name: "Aлекс",
	Email:"alex@yandex.ru",
}
//сериализуем экземпляр в json
resJSON, err := json.Marshal(p)
	if err != nil {
		log.Fatal("unable marshal to json")
	}

	fmt.Printf("Person %v", string(resJSON)) // Man {"Имя":"Alex","Почта":"alex@yandex.ru"}
}
//описание структуры для сериализации
type Person struct {
    Name        string `json:"Имя:"`
    Email       string `json:"Почта:"`
    DateOfBirth time.Time `json:"-"` // - означает, что это поле не будет сериализовано
} 