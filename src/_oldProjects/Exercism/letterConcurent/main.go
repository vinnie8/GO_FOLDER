package letter

import (
	"sync"
)
// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int
//Mutex helps to synchronize acces to memory to avoid deadlocks or other bad things 
var lock = sync.RWMutex{}

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
    var (
    	m = FreqMap{}
        wg sync.WaitGroup
	 )
  //invoking goroutines to count symbols 
	for i:=0;i<len(l);i++{
        wg.Add(1)
        /* If struct is just a map, it's only passed in by reference anyway 
		so you don't have to worry about creating methods that act on pointers
		to avoid copying the entire structure every time.*/
        go counter(l[i],m,&wg)
    }
	//waiting for all goroutines to finish their computation
    wg.Wait()
  	return m
}

func counter(s string, m FreqMap, wg *sync.WaitGroup){
    defer wg.Done()
    
	for _, r:= range s {
        //exclusive access to the map by thread locking
		lock.Lock()
        m[r]++
		lock.Unlock()
	}
	
}

//go test
//go test -bench .
/*
BenchmarkSequentialFrequency-12             5427            227643 ns/op
BenchmarkConcurrentFrequency-12             1274            924789 ns/op
*/

/*
For a good solution, you should see that the concurrent version shows a 
lower number of nano seconds per operation (ns/op) than the sequential version.
*/