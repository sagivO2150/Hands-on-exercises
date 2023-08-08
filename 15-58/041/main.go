package main

import "fmt"

func main() {

	xs := []string{"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)", "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)", "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF) ", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)"}
	// fmt.Printf("the slice - %v\n", xs)
	fmt.Printf("the lengh of the slice is - %v\n", len(xs))
	fmt.Printf("the type is %T\n", xs)

	for i := 0; i < len(xs); i++ {

		fmt.Printf(" taste - %v\n", xs[i])
		// {
		// 	if key < 10 {
		// 		fmt.Printf("icecream taste - %v \t taste - %v\n", key, value)
		// 	} else {
		// 		fmt.Printf("icecream taste - %v \t taste - %v\n", key, value)
		// 	}
		// }
	}

	// fmt.Println(xs[0])
	// fmt.Println(xs[1])
	// fmt.Println(xs[2])
	// fmt.Println(xs[3])
}
