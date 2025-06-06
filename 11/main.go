package main

import (
	"fmt"
)

type impl struct{} //создается тип данных impl

type I interface { //создается тип данных интерфейс
	C()
}

func (*impl) C() {} //создается метод на стракт,
//возвращает С

func A() I { //функция A, ничего не берет, возвращает I-интерфейс
	return nil
}

func B() I { //функция B, ничего не берет, возвращает I-интерфейс
	var ret *impl //переменная-указатель ret на структуру impl
	return ret    //возвращает указание на impl
}

func main() {
	a := A() // объявление переменной, ф-ция возвращает nil
	b := B() // приравнивается указателю на impl-структуру,
	//это не nil
	fmt.Println(a == b) //выдаст false (nil != *impl)
}
