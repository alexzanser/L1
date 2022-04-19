package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func binarySearch(arr[]int, x int) (int) {
	l, m:= 0, 0
	r := len(arr) + 1
	for l < r - 1 {
		m = (l + r) / 2
		if (arr[m] <= x){
			l = m
		} else {
			r = m
		}
	}
	return r
}

func main() {
	var arr []int
	var x int
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < 10; i++ {
		arr = append(arr, int(rand.Int31()%100))
	}
	sort.Ints(arr)
	x = int(rand.Int31()%100)
	fmt.Println(arr)
	fmt.Println(x)
	fmt.Printf("place for x is: %d\n", binarySearch(arr, x))
}
