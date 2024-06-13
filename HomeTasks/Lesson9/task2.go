package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	for i := range s1 {
		fmt.Printf("V1:%d", s1[i])
		fmt.Println()
	OUT:
		for i := range s1 {
			fmt.Printf("\tV2:%d", s1[i])
			fmt.Println()
			for i := range s1 {
				fmt.Printf("\t\tV3:%d", s1[i])
				fmt.Println()
				for i := range s1 {
					fmt.Printf("\t\t\tV4:%d", s1[i])
					fmt.Println()
					if i == 1 {
						break OUT
					}
				}
			}
		}
	}
}
