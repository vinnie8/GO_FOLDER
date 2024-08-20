package main

import(
	"fmt"

 	"github.com/IlshatGadelshin/GO_FOLDER/src/yandexPracticum/sprint2/mathpackage"
)

func main(){

	if sum:=mathpackage.Add(1,2); sum!=3{
		panic(fmt.Sprintf("sum expected to be 3: got %d",sum))
	}

	fmt.Println("Well done!")
}