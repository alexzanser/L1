package main

import (
	"fmt"
	"reflect"
)

/*
func Separate(z interface{}) {
	switch v := z.(type) {
	case int:
		fmt.Printf("I'm int: %d\n", v)
	case string:
		fmt.Printf("I'm string: %s\n", v)
	case bool:
		fmt.Printf("I'm bool: %t\n", v)
	case chan interface{}://как с каналами не понял(
		fmt.Printf("I'm channel: %v\n", v)
	}
}
*/

func GetType(unknown interface{}) {
	fmt.Println(unknown, reflect.TypeOf(unknown))
}

func main() {

	ch1 := make(chan interface{})
	GetType(ch1)
	
	ch2 := make(chan int)
	GetType(ch2)

	GetType("abc")
	GetType(123)
	GetType(true)
}
