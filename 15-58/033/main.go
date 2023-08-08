package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Printf("the big loop ran - %v times\n", i)
		for y := 1; y <= 5; y++ {
			fmt.Printf("the small loop ran %v times\n", y)

		}
		fmt.Println("")
	}

}
