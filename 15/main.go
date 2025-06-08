package main

import (
	"fmt"
)

type myError struct { //создание структуры
	code int //с одним значением типа int
}

func (e myError) Error() string { //определяем метод Error для структуры
	//возвращает строку
	return fmt.Sprintf("code: %d", e.code) //с данными единств. поля
}

func run() error { //функция run возвращает ошибку
	var e *myError //переменная-указатель на myError стракт
	if false {     //если false
		e = &myError{code: 123} //e получает значение указателя
	}
	return e //возвращаем ошибку e
}

func main() {
	err := run()    //err объявляется как результат функции run
	if err != nil { // идем по этой ветке, т.к. указатель не равен nil
		fmt.Println("failed to run, error: ", err) //получим этот вывод
		//ошибка 123 будет
	} else {
		fmt.Println("success")
	}
}
