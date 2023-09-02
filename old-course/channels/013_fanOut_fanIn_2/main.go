package main

import (
	"fmt"
	"sync"
)

var count int

func main() {

	in := gen()

	//fan out
	c1 := factorial(in)
	c2 := factorial(in)
	c3 := factorial(in)
	c4 := factorial(in)
	c5 := factorial(in)

	//fan in
	for n := range merge(c1, c2, c3, c4, c5) {
		count++
		fmt.Println(count, "\t", n)
	}

	// for n := range f {
	// 	fmt.Println(n)
	// }
}

func gen() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
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

func merge(in ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(in))
	for _, i := range in {
		go func(in chan int) {
			defer wg.Done()
			for z := range in {
				out <- z
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

/*
CHALLENGE #1:
-- Change the code above to execute 1,000 factorial computations concurrently and in parallel.
-- Use the "fan out / fan in" pattern

CHALLENGE #2:
WATCH MY SOLUTION BEFORE DOING THIS CHALLENGE #2
-- While running the factorial computations, try to find how much of your resources are being used.
-- Post the percentage of your resources being used to this discussion: https://goo.gl/BxKnOL
*/

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
