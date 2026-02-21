package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	var buf *bytes.Buffer //переменная структуры буффер
	f(buf)                // бросили переменную в функцию и если
	//не равен нулю, пишем строку
	if buf != nil {
		fmt.Printf(buf.String()) //Hello world
	}
	fmt.Println("Main completed") //Main completed
}

func f(out io.Writer) { //функция принимает из райтера
	if out != nil { //если райтер не равен нулю
		_, err := out.Write([]byte("Hello world\n"))
		//пустая переменная и ошибка присваивает результат
		//вывода  метода райт от переменной out
		//если ошибка
		if err != nil {
			log.Fatal(err)
			//логируем фаталом
		}
	}
}
