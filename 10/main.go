package main

import (
	"fmt"
)

func main() {
	var s *string         //указатель s на строковую переменную
	fmt.Println(s == nil) //сама переменная нулевая, -> true
	var i interface{}     //пустой интерфейс i
	fmt.Println(i == nil) //сравнение с нулем true
	i = s                 //интерфейс становится равен s (указателю)
	fmt.Println(i == nil) //false
}
