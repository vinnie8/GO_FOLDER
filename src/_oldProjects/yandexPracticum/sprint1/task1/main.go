package main

import "fmt"
/*Дополните пример объявлениями переменных ver, id, pi так, 
чтобы программа выводила ver = v0.0.1 id = 0 pi = 3.1415.*/

func main() {
    // определите переменные ver, id, pi
	var(
		ver string="v.0.0.1"
		id int
		pi float32=3.1415
	)

    fmt.Println("ver =", ver, "id =", id, "pi =", pi)
}