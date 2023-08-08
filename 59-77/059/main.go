package main

import "fmt"

func main() {
	total := foo([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(total)

	total2 := bar([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(total2)
}

func foo(nums ...int) int {
	t := 0
	for _, v := range nums {
		t += v
	}
	return t
}

func bar(nums []int) int {
	t := 0
	for _, v := range nums {
		t += v
	}
	return t
}

/*
Hands-on exercise #59 - variadic func
● create a func with the identifier foo that
○ takes in a variadic parameter of type int
○ pass in a value of type []int into your func (unfurl the []int)
○ returns the sum of all values of type int passed in
● create a func with the identifier bar that
○ takes in a parameter of type []int
○ returns the sum of all values of type int passed in
*/
