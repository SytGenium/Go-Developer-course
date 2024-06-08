package main

import (
	"fmt"
)

type contact struct {
	adress string
	phone  string
}
type user struct {
	ID   int
	Name string
	contact
}
type employee struct {
	ID   int
	Name string
	contact
}

	U1 := contract{
		ID:     1,
		Name: "Юзер1",
		Date:   "2024-01-31",
		contact: contact{
			adress: "Орел",
			phone: "9045674",
			}
		}
E1 := contract{
ID:     1,
Name: "Юзер1",
Date:   "2024-01-31",
contact: contact{
adress: "Орел",
phone: "9045674",
}
}

func main() {
	fmt.Printf("Договор № %s от %s", W.Number, W.Date)
}
