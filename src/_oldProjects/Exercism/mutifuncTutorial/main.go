package main

//https://www.digitalocean.com/community/tutorials/how-to-run-multiple-functions-concurrently-in-go
import (
	"fmt"
	"sync"
)

func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
	for i := 1; i <= total; i++ {
		fmt.Printf("sending %d to channel\n", i)
		ch <- i
	}
}

func printNumbers(idx int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range ch {
		fmt.Printf("%d: read %d from channel\n", idx, num)
	}
}

func main() {
	//adding a waitgroup
	var wg sync.WaitGroup
	numberChan := make(chan int)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go printNumbers(i, numberChan, &wg)
	}

	generateNumbers(5, numberChan, &wg)

	close(numberChan)

	fmt.Println("Waiting for goroutines to finish...")
	wg.Wait()
	fmt.Println("Done!")
}
