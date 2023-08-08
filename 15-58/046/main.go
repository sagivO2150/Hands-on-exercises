package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("before editing -", x)
	// var y []int

	// for _, v := range x {
	// 	if v <= 44 || v >= 48 {
	// 		y = append(y, v)
	// 	} else {
	// 		continue
	// 	}
	// }

	// x = y

	x = append(x[:3], x[6:]...)
	fmt.Println("after editing -", x)

}
