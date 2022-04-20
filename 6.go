package main

/*
//Закрытие канала

import (
	"fmt"
	"time"
)

func Worker(ch chan struct{}) {
	for range ch {}
	fmt.Println("Worker died")
}

func main() {
	ch := make(chan struct{})

	go Worker(ch)

	time.Sleep(time.Second * 3)
	close(ch)//Закрываем канал, воркер понимает, что канал закрылся и умирает
	time.Sleep(time.Second * 3)

	fmt.Println("Main died")
}
*/

/*
//С помощью SIGINT

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func Worker(shutdown chan os.Signal, wg *sync.WaitGroup) {
	defer wg.Done()
	signal.Notify(shutdown, syscall.SIGINT)//
	LOOP:
		for  {
			select {
			case <- shutdown:
				break LOOP
			}
		}
	fmt.Println("Worker died")
}

func main() {
	var  wg sync.WaitGroup
	shutdown := make(chan os.Signal, 1)

	signal.Ignore(syscall.SIGINT)

	wg.Add(1)
	go Worker(shutdown, &wg)
	wg.Wait()
}
*/


/*
//С помощью контекста
import (
	"context"
	"fmt"
	"time"
)
func Worker(ctx context.Context) {
	LOOP:
		for {
			select {
			case <- ctx.Done():
				break LOOP
			}
		}
	fmt.Println("Worker died")
}

func main() {
	ctx := context.TODO()
	ctx, cancel := context.WithCancel(ctx)
	go Worker(ctx)

	time.Sleep(time.Second * 3)
	cancel()//завершаем контекст
	time.Sleep(time.Second * 3)

	fmt.Println("Main died")
}
*/
