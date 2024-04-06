package main

import "fmt"

func main() {
	a := 16
	b := 3
	c := a / b
	fmt.Printf("Результат: %v, остаток от деления: %v, тип результата: %T", c, a%b, c)
}
