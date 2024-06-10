package main

import (
	"fmt"
)

func main() {
	m1 := [4]string{"яблоко ", "груша ", "помидор ", "абрикос"}
	fmt.Println(m1)
	m1[2] = "персик "
	fmt.Println(m1)
}
