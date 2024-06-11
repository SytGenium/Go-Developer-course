package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := append(s1[:3], s1[4:]...)
	fmt.Println(s2)
}
