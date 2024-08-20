/*Есть два модуля, которые расположены относительно
 друг друга следующим образом:
 ├── main
│   ├── go.mod
│   └── main.go
└── somemodule
    ├── somepackage
    │   └── file.go
    └── go.mod 
	
	Напишите файл main/go.mod так, 
	чтобы код отработал без ошибок.*/

package main

import (
   
    "somemodule/somepackage"
)

func main() {
    somepackage.Func()
}