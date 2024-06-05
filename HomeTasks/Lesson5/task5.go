package main

import "fmt"

func main() {
	var d int
	d = 50
	fmt.Println(d)
	fmt.Println(change(d))
}

func change(a int) int {
	var b *int
	b = &a
	*b = 100
	return a
}
