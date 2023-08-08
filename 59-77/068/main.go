package main

import "fmt"

func main() {
	full := func(first, last string) []string {
		fullName := make([]string, 3)
		fullName[1] = first
		fullName[2] = last

		return fullName

	}

	fmt.Println(full("sagiv", "oron"))

	x()

	y := func() {
		for i := 0; i <= 10; i++ {
			fmt.Println(i)
		}
	}

	y()
}

var x = func() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

/*
Hands-on exercise #68 - anonymous func
● Build and use an anonymous func
● anonymous func / func literal / lambda func
an anonymous func, aka a function literal or a lambda function, is a way to define a function without giving it a name. Instead, you can directly declare and define a function inline wherever it is needed. Anonymous funcs are commonly used when you want to pass a function as an argument to another function or when you need to define a short-lived function for a specific task.
*/
