package main

import "fmt"

type person struct {
	first string
}

func changeIt(p person) person {
	p.first = "kaki"
	return p
}

func changeItP(p *person) {
	p.first = "pipi"
}

func main() {
	p1 := person{
		first: "sagiv",
	}

	fmt.Println("the original name was ", p1.first)
	p1 = changeIt(p1)
	fmt.Println("the new name is ", p1.first)
	changeItP(&p1)
	fmt.Println("and the new new name is ", p1.first)
}

/*
Hands-on exercise #77 - value & pointer semantics
Create two functions to change a field in a struct called "first" of TYPE string:
● One function will use VALUE SEMANTICS
○ this function will return some VALUE of some TYPE
● The other function will use POINTER SEMANTICS ○ this function will return nothing
● don't use methods
*/
