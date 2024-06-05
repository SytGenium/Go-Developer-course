package main

import "fmt"

func main() {
	var d int
	d = 50
	fmt.Println(d)
	fmt.Println(pr(d))
}

func pr(a int) int {
	var b *int
	b = &a
	*b = 100
	return a
}
