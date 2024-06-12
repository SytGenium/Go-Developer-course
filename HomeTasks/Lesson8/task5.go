package main

import "fmt"

func main() {
	s1 := map[string]int{"слон": 3, "бегемот": 0, "носорог": 5, "лев": 1}
	s1["бегемот"] = 2
	fmt.Println(s1)
}
