package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	c1 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "toyota",
		model: "98",
		doors: 4,
		color: "black",
	}

	c2 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "ferrari",
		model: "04",
		doors: 2,
		color: "red",
	}

	fmt.Println(c1, c2)

	fmt.Printf("the maker of the cars are %v and %v\n", c1.make, c2.make)

}
