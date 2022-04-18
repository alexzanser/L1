package main

import "fmt"

type Human struct {}

func (h Human) Greeting() {
	fmt.Println("I am Human!")
}

type Action struct {
	Human
}

func (a Action) Greet() {
	a.Greeting()
	fmt.Println("I am acting")
}


func main() {
	var act Action
	act.Greet()
}
