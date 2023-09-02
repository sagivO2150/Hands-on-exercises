package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("my name is ", p.first, p.last, "and my age is ", p.age)
}

type human interface {
	speak()
}

func saySomthing(h human) {
	h.speak()
}

func main() {
	p1 := &person{
		first: "sagiv",
		last:  "oron",
		age:   27,
	}

	p2 := person{
		first: "gal",
		last:  "oron",
		age:   25,
	}

	saySomthing(p1)
	// saySomthing(p2)
	p2.speak()

}

/*
Hands-on exercise #2
This exercise will reinforce our understanding of method sets:
● create a type person struct
● attach a method speak to type person using a pointer receiver
○ *person
● create a type human interface
○ to implicitly implement the interface, a human must have the speak method
● create func “saySomething”
○ have it take in a human as a parameter
○ have it call the speak method
● show the following in your code
○ you CAN pass a value of type *person into saySomething
○ you CANNOT pass a value of type person into saySomething
*/
