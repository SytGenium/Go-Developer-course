package main

import "fmt"

func main() {
	a := 16
	b := 3
	c := a / b
	a = a - c*b
	fmt.Println(c)
	fmt.Println(a)
	fmt.Printf("Результат: %v, остаток от деления: %v, тип результата: %T", c, a, c)
}
