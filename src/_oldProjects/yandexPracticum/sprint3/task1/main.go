/*
Реализуйте тип Stopwatch, который будет сохранять время каждой промежуточной фиксации секундомера и выдавать результаты относительно общего времени старта.
Тип должен обладать следующими методами:
Start() — запустить/сбросить секундомер;
SaveSplit() — сохранить промежуточное время;
GetResults() []time.Duration — вернуть текущие результаты.
*/

package main

import (
	"fmt"
	"time"
)

//Пример работы:
func main() {
	sw := Stopwatch{}
	sw.Start()

	time.Sleep(1 * time.Second)
	sw.SaveSplit()

	time.Sleep(500 * time.Millisecond)
	sw.SaveSplit()

	time.Sleep(300 * time.Millisecond)
	sw.SaveSplit()

	fmt.Println(sw.GetResults())

}

//тип, содержащий поля с временем старта и результата
type Stopwatch struct {
	startTimer time.Time
	result     []time.Duration
}

//создает точку отсчета/старта
func (s *Stopwatch) Start() {
	s.startTimer = time.Now()
}

//сохраняет пройденное время на момент вызова
func (s *Stopwatch) SaveSplit() {
	t := time.Now()
	s.result = append(s.result, t.Sub(s.startTimer))
}

//выводит результат работы
func (s Stopwatch) GetResults() []time.Duration {
	return s.result
}

//gofmt -w main.go
