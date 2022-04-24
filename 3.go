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

	start := time.Now().Unix()
	sum := 0
	arr := []int{2, 4, 6, 8, 10}

	for _, val := range arr {
		wg.Add(1)
		go SimulateCalc(val, &wg, &mu, &sum)//каждая функция работает 5 секунд
	}
	wg.Wait()
	fmt.Println(sum)
	//вычисляем параллельно поэтому на все уходит 5 секунд
	fmt.Println("Seconds passed since programm start :", time.Now().Unix() - start)
}
