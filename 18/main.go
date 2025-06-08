package main

import "fmt"

func main() {
	var s *string                //переменная s - указатель строковый
	fmt.Println( /*&*/ s == nil) //false / true |получается, что
	//указатель не равен нулю, а переменная сама равна нулю
	var i interface{}
	fmt.Println(i == nil) //true
	i = s
	fmt.Println(i == nil) //false
}

// true
// true
// false
