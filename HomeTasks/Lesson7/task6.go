package main

import "fmt"

func main() {
	s1 := make([]int, 0, 10)
	s1 = append(s1, 4, 1, 8, 9)
	fmt.Println(s1)
}
