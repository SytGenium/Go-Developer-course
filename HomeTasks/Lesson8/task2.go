package main

import (
	"fmt"
)

var animals = [3]string{"слон", "бегемот", "осьминог"}
var s1 = map[string]int{"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1}

func main() {
	f := 2
	vwvod(f)
}

func vwvod(a int) {
	if a < 0 {
		return
	}
	m, ok := s1[animals[a]]
	fmt.Printf(":Животное: %s, количество: %d (есть в карте: %v)", animals[a], m, ok)
	fmt.Println()
	a--
	vwvod(a)
}
