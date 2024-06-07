package main

import (
	"fmt"
)

type contract struct {
	ID     int
	Number string
	Date   string
}

func main() {
	var W = contract{
		ID:     1,
		Number: "#000A\\n101",
		Date:   "2024-01-31",
	}
	fmt.Printf("{ID:%d Number%s Date:%s", W.ID, W.Number, W.Date)
}
