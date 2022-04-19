package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int64
} 

func main() {
	var counter Counter
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func () {
			defer wg.Done()
			atomic.AddInt64(&counter.value, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter.value)
}
