package main

import "fmt"

func main() {
	var s string

	fmt.Println("Enter string value:")
	fmt.Scan(&s)

	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		fmt.Printf("%c", runes[i])
	}
	fmt.Println()
}
