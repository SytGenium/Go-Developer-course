package main

import "fmt"

func main() {
	type animal struct{}
	emptyanimal := animal{}

	s1 := map[string]animal{}
	s1["слон"] = emptyanimal
	s1["бегемот"] = emptyanimal
	s1["носорог"] = emptyanimal
	s1["лев"] = emptyanimal
	fmt.Println(s1)
}
