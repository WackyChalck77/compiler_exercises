package main

import "fmt"

type Person struct { //вводим стракт с одним полем
	Name string
}

func changeName(person *Person) { //функция принимает этот стракт (не копию)
	// person = &Person{ //принимаем значение по указателю
	// 	Name: "Alice", //переназначаем сам указатель, чтобы он указывал на
	//новый объект Person с другим полем Name Alice
	// }
	//здесь нужно сделать разыменовывание указателя
	person.Name = "Alice"
}

func main() {
	person := &Person{ //объявляем переменную указатель типа Person
		Name: "Bob", //даем поле
	}
	fmt.Println(person.Name) //печатаем Bob
	changeName(person)       // меняем на Alice
	fmt.Println(person.Name) //печатаем Alice
}
