package main

import "fmt"

func squerePrint(f func(int) int, n int) string {
	x := f(n)
	return fmt.Sprintf("the squre of %v is - %v\n", n, x)
}

func squareThis(n int) int {
	return n * n
}

func main() {
	fmt.Println(squerePrint(squareThis, 5))
}
