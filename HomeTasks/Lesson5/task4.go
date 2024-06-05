package main

import "fmt"

var a int
var b int
var c int
var d int

func main() {
	pr(a)
	pr(b)
	pr(c)
	pr(d)
}

func pr(a int) {
	fmt.Println(a, &a)
}
