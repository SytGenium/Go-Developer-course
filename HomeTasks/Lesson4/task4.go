package main

import "fmt"

func main() {
	var f func() string
	f = func() string {
		return "Hello, GO!"
	}

	hello(f)
}

func hello(v func() string) {
	fmt.Println(v())
}
