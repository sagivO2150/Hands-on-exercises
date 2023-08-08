package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}

	for k, v := range xi {
		fmt.Printf("the value in place numeber - %v\t is - %v\n", k, v)
	}
}
