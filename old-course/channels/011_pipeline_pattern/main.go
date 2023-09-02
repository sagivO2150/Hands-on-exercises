package main

import (
	"fmt"
)

func main() {
	in := gen()

	f := factorial(in)

	for z := range f {
		fmt.Println(z)
	}
}

func gen() chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < 10; i++ {
			for z := 3; z < 13; z++ {
				// fmt.Println(z)
				out <- z
			}
		}
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			// fmt.Println(i)
			res := fact(i)
			out <- res
		}
	}()
	return out

}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
