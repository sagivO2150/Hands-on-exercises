package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "sagiv"
}
