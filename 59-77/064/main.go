package main

import "fmt"

func main() {
	fmt.Printf("%T\n", Add)
	fmt.Printf("%T\n", Subtract)
	fmt.Printf("%T\n", doMath)

	x := doMath(42, 16, Add)
	fmt.Println(x)

	y := doMath(42, 16, Subtract)
	fmt.Println(y)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
func Add(a int, b int) int {
	return a + b
}
func Subtract(a int, b int) int {
	return a - b
}
