package main

import (
	"fmt"
)

func main() {
	xi := []int{42, 43, 44, 45, 46, 47}
	for k, v := range xi {
		fmt.Printf("the key is - %v and the value is - %v\n", k, v)
	}

	m := map[string]int{
		"sagiv": 27,
		"gal":   25,
	}

	for k, v := range m {
		fmt.Printf("his name is - %v and his age is - %v\n", k, v)
	}
}
