package main

import "fmt"

type square int

func main() {
	var s square
	s = 34
	s = s + 10
	square_out(s)
}

func square_out(f square) {
	fmt.Println(f, "м2")
}
