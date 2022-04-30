package main

import (
	"fmt"
	"sync"
)

func main() {
	
	arr := []int{1, 2, 3, 4, 5}
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func () {
		defer close(ch2)
		for val := range ch1 {
			ch2 <- val * val
		}
	} ()
	
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for val := range ch2 {
			fmt.Println(val)
		}
	}()

	for _, val := range arr {
		ch1 <- val
	}

	close(ch1)
	wg.Wait()
}
