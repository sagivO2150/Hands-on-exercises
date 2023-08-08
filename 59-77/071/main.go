package main

import "fmt"

type test func(int) bool

func even(s int) bool {
	return s%2 == 0
}

func odd(s int) bool {
	return s%2 != 0
}

func sumEvenOrOdd(n []int, f test) int {
	sum := 0
	for _, v := range n {
		if f(v) {
			sum += v
		}
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sumEven := sumEvenOrOdd(nums, even)
	fmt.Println("the sum of the even numbers is -", sumEven)
	sumOdd := sumEvenOrOdd(nums, odd)
	fmt.Println("the sum of the even numbers is -", sumOdd)
}

/*
 Hands-on exercise #71 - callback
A “callback” is when we pass a func into a func as an argument. For this exercise,
● pass a func into a func as an argument
○ func square(n int) int
○ printSquare(f func(int)int, int) string
*/
