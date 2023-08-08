package main

import "fmt"

func main() {

	xi := [5]int{4, 3, 2, 1, 0}

	for key, value := range xi {
		fmt.Printf("the value of position number - %v in the array xi is - %v and it's type is - %T\n", key, value, value)
	}
}
