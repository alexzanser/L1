package main

import "fmt"

func main() {
	x, y := 1, 2
	fmt.Printf("Pre:  x=%d; y=%d\n", x, y)
	x, y = y, x
	fmt.Printf("Post: x=%d; y=%d\n", x, y)
}