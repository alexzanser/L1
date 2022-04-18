package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	alive , _ := strconv.Atoi(os.Getenv("ALIVE"))

	start := time.Now().Unix()
	ch := make(chan int, 1)

	s := 0
	g := 0
	for time.Now().Unix() - start < int64(alive) {
		ch <- s
		fmt.Printf("Send : %d\n", s)
		g = <- ch
		fmt.Printf("Get  : %d\n", g)
		time.Sleep(time.Second)
		s++
	}
}
