package main

import (
	"fmt"
	"unsafe"
)

type some1 struct { //создается структура
	a bool
	b int32
	c string
}

type some2 struct { //создается структура
	b int32
	c string
	a bool
}

// не происходит ничего
func main() {
	a := unsafe.Sizeof(some1{}) //создаем пустое значение чтобы
	//узнать объем памяти
	fmt.Println(a)

	b := unsafe.Sizeof(some2{}) //создаем пустое значение чтобы
	//узнать объем памяти
	fmt.Println(b)
}
