package main

import (
	"sync"
	"time"
	"fmt"
)

func SimulateCalc(x int, wg *sync.WaitGroup, mu *sync.Mutex, sum *int) {
	defer wg.Done()

	time.Sleep(time.Second * 5)
	mu.Lock()
	*sum += x * x
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	sum := 0
	arr := []int{2, 4, 6, 8, 10}

	for _, val := range arr {
		wg.Add(1)
		go SimulateCalc(val, &wg, &mu, &sum)//Каждая функция работает 5 секунд
	}
	wg.Wait()
	fmt.Println(sum)
}
