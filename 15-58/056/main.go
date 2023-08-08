package main

import "fmt"

func main() {
	p1 := struct {
		first     string
		frinds    map[string]int
		favDrinks []string
	}{
		first: "sagiv",
		frinds: map[string]int{
			"garten": 28,
			"gal":    25,
			"daniel": 900,
		},
		favDrinks: []string{
			"pepsi",
			"soda",
			"vodka",
		},
	}

	// fmt.Println(p1)

	fmt.Printf("hello, my name is %v, my friends are - \n\n", p1.first)

	for k, v := range p1.frinds {
		fmt.Printf("%v, his age is %v\n", k, v)
	}

	fmt.Println()
	fmt.Println("some of my favorite drinks are -")

	for _, v := range p1.favDrinks {
		fmt.Println(v)
	}

}

/*
Create and use an anonymous struct with these fields:
● first string
● friends map[string]int
● favDrinks []string
Print things.
*/
