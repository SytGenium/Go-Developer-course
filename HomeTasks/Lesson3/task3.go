package main

import "fmt"

const myGlobalConst = 679085

func main() {
	const myGlobalConst = 1
	fmt.Printf("Моя первая константа %d!", myGlobalConst)
}
