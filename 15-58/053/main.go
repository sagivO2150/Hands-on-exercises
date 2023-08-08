package main

import "fmt"

type person struct {
	fist            string
	last            string
	icecreamFlavors []string
}

func main() {

	p1 := person{
		fist:            "sagiv",
		last:            "oron",
		icecreamFlavors: []string{"mango", "lemon", "banana"},
	}
	p2 := person{
		fist:            "gal",
		last:            "ohayon",
		icecreamFlavors: []string{"cherry", "chocolate", "blue barries"},
	}

	fmt.Printf("this is: %v %v and his favorite icecream flavors are - \n", p1.fist, p1.last)
	for _, v := range p1.icecreamFlavors {
		fmt.Printf("%v\n", v)
	}

	fmt.Printf("and this is: %v %v and his favorite icecream flavors are - \n", p1.fist, p1.last)
	for _, v := range p2.icecreamFlavors {
		fmt.Printf("%v \n", v)
	}

	fmt.Println("------------------------------------------")

	m := make(map[string]person)
	m[p1.last] = p1
	m[p2.last] = p2

	for _, v := range m {
		// fmt.Printf("%#v %#v", k, v)
		// fmt.Printf("%T %T\n", k, v)
		fmt.Printf("the favorit ice cream flavors of %v %v are -\n", v.fist, v.last)
		for _, v1 := range v.icecreamFlavors {
			fmt.Println(v1)
		}
	}

}
