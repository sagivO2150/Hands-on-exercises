package main

import "fmt"

func main() {

	i := 0

	for {

		if i > 9 {
			fmt.Printf("this loop sttoped after %v times\n", i)
			break
		}
		fmt.Printf("this loop ran %v times\n", i)
		i++
	}
}
