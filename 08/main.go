package main

import (
	"fmt"
)

type Animal interface{} //объявление интерфейса животное
type Dog struct{}       //объявление структуры собака

func IsNil(i interface{}) bool { //функция, которая берет интерфейс и
	//выдает булево значение
	return i == nil //возвращаем сравнение с нулём интерфейса
}

func main() {
	var a Animal //объявляем переменную а типа животное
	var d *Dog   //объявляем указатель на структуру собака

	fmt.Println(IsNil(a)) //выдаст true
	fmt.Println(IsNil(d)) //выдаст true  False, т.к. это
	//указатель, он не нулевой
}
