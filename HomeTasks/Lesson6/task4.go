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

var u = user{
	ID:   1,
	Name: "Юзер1",
	contact: contact{
		adress: "Орел",
		phone:  "9045674",
	},
}
var e = employee{
	ID:   1,
	Name: "Шеф1",
	contact: contact{
		adress: "Москва",
		phone:  "982098208024",
	},
}

func main() {
	fmt.Println(u.contact.adress, u.contact.phone, e.contact.adress, e.contact.phone)
}
