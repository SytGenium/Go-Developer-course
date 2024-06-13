package main

import (
	"fmt"
)

func main() {
	q1 := []string{"груша", "яблоко", "бегемот", "тыква", "огурец", "помидор"}
	for i := range q1 {
		checkFood(q1[i])
	}
}

func checkFood(a string) {
	switch a {
	case "груша":
		fallthrough
	case "яблоко":
		fallthrough
	case "апельсин":
		fmt.Println("Это фрукт")
	case "тыква":
		fallthrough
	case "огурец":
		fallthrough
	case "помидор":
		fmt.Println("Это овощ")
	default:
		fmt.Println("что-то странное…")
	}

}
