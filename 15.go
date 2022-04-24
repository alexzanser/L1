package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc(c string) {
	v := strings.Repeat(c, 1 << 10)
	justString = v[:100]
}

func someFuncFixed(c string) {
	v := strings.Repeat(c, 1 << 25)

	bufSlice := make([]rune, 100)
	copy(bufSlice, []rune(v)[:100])

	justString = string(bufSlice)
}

/*в случае если символ занимает более одного байта, то срез по слайсу байт дает не тот результат,
что ожидается

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}


justSring глобальная переменная, поэтому выделенная под hugeString память не будет очищаться
*/

func main() {
  someFunc("異")
  fmt.Println("Incorrect: ", justString)//некорректный вывод

  someFuncFixed("異")
  fmt.Println("Correct: ", justString)//корректный вывод

}
