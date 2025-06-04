package main

import (
	"fmt"
)

func main() {
	digits := []int{1, 2, 3, 4, 5} //инициализируем слайс
	for _, d := range digits {
		defer fmt.Println(d) //выводим все элементы с конца|верно!
	}
}
