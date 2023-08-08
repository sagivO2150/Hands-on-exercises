package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// x1 := xi[:5]
	// fmt.Println(x1)
	// x2 := xi[5:]
	// fmt.Println(x2)
	// x3 := xi[2:7]
	// fmt.Println(x3)
	// x4 := xi[1:6]
	// fmt.Println(x4)

	fmt.Println(xi[:5])
	fmt.Println(xi[5:])
	fmt.Println(xi[2:7])
	fmt.Println(xi[1:6])

	// for _, v := range xi {
	// 	fmt.Printf("the value is %v and the type is %T\n", v, v)
	// }

}
