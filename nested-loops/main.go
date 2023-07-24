package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 6; i++ {
		fmt.Printf("the outter loop ran - %v \n", i)
		for y := 1; y < 6; y++ {
			fmt.Printf("\t\t\t the inner loop ran - %v times\n", y)
		}
	}
}
