package main

import "fmt"

var a string
var b *string

func main() {
	a = "Hello"
	fmt.Println(a)
	b = &a
	*b = "GO!!!"
	fmt.Println(a)
}
