package main

import "fmt"

func main() {
	fmt.Println("Hello, GO!")
	defer fmt.Println("Завершение работы.")
}
