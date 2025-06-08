package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() { //отложенная анонимная функция
		if err := recover(); err != nil {
			//отлов паники с recover вне горутины с паникой
			fmt.Println(err) //печатаем ошибку
		}
	}()

	go func() { //горутина!
		panic(123) //паника с ошибкой 123
	}()

	time.Sleep(time.Second) //засыпаем на час? Чё за пздц? Hour
}
