package main

import "fmt"

var a string

func main() {
	a = "Hello"
	fmt.Println(a, &a)
}
