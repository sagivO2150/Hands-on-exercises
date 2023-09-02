package main

import (
	"fmt"
)

type customErr struct {
	message string
	code    int
}

func (ce customErr) Error() string {
	return fmt.Sprintf("%s and the code is %v", ce.message, ce.code)
}

func main() {
	c1 := customErr{
		message: "this is sagiv's custom error",
		code:    400,
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}

/*
Hands-on exercise #3
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. If you need a hint, here is one.
*/
