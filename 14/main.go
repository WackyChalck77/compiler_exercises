package main

import (
	"fmt"
)

type Example struct { //создаем структуру
	Value string // где одно поле строковое
}

type MyInterface interface{} //создаем интерфейс

func example1() MyInterface { //метод для интерфейса
	var e *Example //объявляется переменная-указатель структуры Example
	return e       //возвращаем её значение
}

func example2() MyInterface { //метод интерфейса
	return nil //возвращает nil
}

func main() {
	fmt.Println(example1() == example2()) //не ноль vs ноль
	// false
}
