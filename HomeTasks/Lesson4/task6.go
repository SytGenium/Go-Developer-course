package main

import "fmt"

func main() {
	recurs(23, 0, 1)
}
func recurs(a int, b int, c int) {
	a--
	fmt.Println(b)

	if a == 0 {
		return
	}
	recurs(a, c, b+c)
}
