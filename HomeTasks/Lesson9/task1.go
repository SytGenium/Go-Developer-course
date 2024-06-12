package main

import (
	"fmt"
)

func main() {
	q1 := [10]string{"яблоки", "банан", "сливы",
		"персики", "лимоны", "черешня",
		"груши", "рябина", "кизил", "терен"}
	for i := 0; i < 10; i++ {
		m, ok := fruitMarket(q1[i])
		if ok == true && m > 0 {
			fmt.Printf("%s осталось %d штук", q1[i], m)
			fmt.Println()
		} else if ok == true && m == 0 {
			fmt.Printf("%s закончились", q1[i])
			fmt.Println()
		} else {
			fmt.Printf("%s не оказалось в магазине", q1[i])
			fmt.Println()
		}
	}
}

func fruitMarket(a string) (int, bool) {
	s1 := map[string]int{"апельсин": 5, "яблоки": 3, "сливы": 1, "груши": 0}
	m, ok := s1[a]
	return m, ok

}
