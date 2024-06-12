package main

import "fmt"

func main() {
	animals := [3]string{"слон", "бегемот", "осьминог"}
	s1 := map[string]int{"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1}

	m, ok := s1[animals[0]]
	fmt.Printf(":Животное: %s, количество: %d (есть в карте: %v)", animals[0], m, ok)
	fmt.Println()
	m1, ok := s1[animals[1]]
	fmt.Printf(":Животное: %s, количество: %d (есть в карте: %v)", animals[1], m1, ok)
	fmt.Println()
	m2, ok := s1[animals[2]]
	fmt.Printf(":Животное: %s, количество: %d (есть в карте: %v)", animals[2], m2, ok)
}
