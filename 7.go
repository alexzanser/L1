package main

import (
	"fmt"
	"sync"
	"time"
)

func AddRecord(carPrices  map[string]int, mu *sync.Mutex, wg *sync.WaitGroup, key string, value int) {
	defer wg.Done()
	time.Sleep(time.Second * 5)

	mu.Lock()
	carPrices[key] = value
	mu.Unlock()
}


func main() {
	carPrices := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	start := time.Now().Unix()
	keys := []string{"Shevrolet", "Opel", "Ford"}
	values := []int{10000, 5000, 7000}

	for i := 0; i < len(keys); i++ {
		wg.Add(1)
		go AddRecord(carPrices, &mu, &wg, keys[i], values[i])//на каждую запись по 5 секунд
	}
	wg.Wait()

	for k, v:= range carPrices {
		fmt.Println(k, v)
	}
	//пишем параллельно поэтому на все уходит 5 секунд
	fmt.Println("Seconds passed since programm start :", time.Now().Unix() - start)
}
