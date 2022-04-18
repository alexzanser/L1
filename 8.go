package main

import "fmt"


func SetBit1(x int64, i int) int64 {
	x = x | (1 << (i));
	return x
}

func SetBit0(x int64, i int) int64 {
	x = x & ^(1<<i)
	return x
}

func main() {
	var x int64 = 64

	fmt.Println(SetBit1(x, 5))//ставим 5 бит в 1, получаем 96
	fmt.Println(SetBit0(x, 5))//возвращаем обратно
	
}