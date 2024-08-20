/*
Отформатируйте утилитой gofmt. Вот пример вызова: gofmt -w ..
Проверьте линтером go vet и исправьте все найденные проблемы.
Добавьте для описания типа и функций комментарии, которые обработает godoc.
Запустите godoc и проверьте сгенерированную документацию в браузере.
*/

package arrint

import (
	"fmt"
	"strings"
)
// ArrInt описывает слайс целых чисел типа int.
type ArrInt []int

//метод складывает элементы двух слайс
func Add(a, b ArrInt) ArrInt {
	length := len(a)
	if length-len(b) > 0 {
		length = len(b)
	}
	//новый слайс длина которого равна самому короткому слайсу из полученных
	c := make(ArrInt, length)
	for i := 0; i < length; i++ {
		c[i] = a[i] + b[i]
	}
	return c
}

// Метод String преобразует ArrInt в строку и возвращает её.
func (a ArrInt) String() string {
	out := make([]string,
		len(a))

	for i, v := range a {
		out[i] = fmt.Sprintf("%d", v)
	}
	return fmt.Sprintf(`[%s]`, strings.Join(out, ` `))
}


//gofmt -w main go
//go vet main.go
//go get -v  golang.org/x/tools/cmd/godoc