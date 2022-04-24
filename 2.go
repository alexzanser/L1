package main

import (
	"sync"
	"time"
	"fmt"
)

func SimulateCalc(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 5)
	fmt.Println(x * x)
}

func main() {
	var wg sync.WaitGroup

	start := time.Now().Unix()
	arr := []int{2, 4, 6, 8, 10}

	for _, val := range arr {
		wg.Add(1)
		go SimulateCalc(val, &wg)////каждая функция работает 5 секунд
	}
	wg.Wait()
	//вычисляем параллельно поэтому на все уходит 5 секунд
	fmt.Println("Seconds passed since programm start :", time.Now().Unix() - start)
}