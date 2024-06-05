package main

import "fmt"

func main() {
	var a *string
	f := "10"
	a = &f
	fmt.Println(a)
}
