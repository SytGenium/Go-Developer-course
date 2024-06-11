package main

import "fmt"

func main() {
	s1 := make([]int, 0, 10)
	fmt.Println("s1 длина :", len(s1), "ёмкость:", cap(s1))
}
