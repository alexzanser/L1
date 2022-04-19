package main

import (
	"fmt"
	"strings"
	"time"
	"unsafe"
)

var justString string

func someFunc() {
	v := strings.Repeat("Я", 1 << 30)
	justString = v[:100]
}

func main() {
  fmt.Println(unsafe.Sizeof(justString))
  someFunc()
  fmt.Println(justString)
  time.Sleep(time.Second * 3)
  fmt.Println(justString)
}
