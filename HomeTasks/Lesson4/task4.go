package main

import "fmt"

func hello() func() string {
	return func() string {
		return "Hello, GO!"
	}
}

func main() {
	f := hello()
	fmt.Println(f())
}
