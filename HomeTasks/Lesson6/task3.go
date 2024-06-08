package main

import (
	"fmt"
)

type contract struct {
	ID     int
	Number string
	Date   string
}

var W = contract{
	ID:     1,
	Number: "#000A\\n101",
	Date:   "2024-01-31",
}

func main() {
	fmt.Printf("Договор № %s от %s", W.Number, W.Date)
}
