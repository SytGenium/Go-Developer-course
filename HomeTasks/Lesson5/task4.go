package main

import "fmt"

var a float32
var b float32
var c float32
var d float32

func main() {
	pr(a)
	pr(b)
	pr(c)
	pr(d)
}

func pr(a float32) {
	fmt.Println(a, &a)
}
