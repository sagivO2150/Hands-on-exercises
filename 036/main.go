package main

import "fmt"

func main() {
	m := map[string]int{"James": 42,
		"Moneypenny": 32}

	for k, v := range m {
		fmt.Println("")
		fmt.Printf("key - %v\nvalue - %v\n", k, v)
		fmt.Println("")
	}
	age := m["Q"]
	fmt.Println(age)
	if v, ok := m["Q"]; !ok {
		fmt.Printf("the value in 'Q' is - %v \n", v)
	}
}
