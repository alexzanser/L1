package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)
func Translator(ch chan  interface{}) {
	i := 0
	for {
		ch <- i
		time.Sleep(time.Millisecond * 100)
		i++
	}
}

func Worker(id int, ch chan  interface{}) {
	for val := range ch {
		fmt.Printf("I'm worker : %d I'm working : %v\n", id, val)
	}
}

func main() {
	ch := make(chan interface{})
	shutdown := make(chan os.Signal, 1)
	n, _ := strconv.Atoi(os.Getenv("WORKERSN"))

	signal.Notify(shutdown, syscall.SIGINT)

	go Translator(ch)

	for i := 0; i < n ; i++ {
		go Worker(i, ch)
	}

	<-shutdown
	close(ch)
}
