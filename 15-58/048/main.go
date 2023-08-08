package main

import "fmt"

func main() {
	xs := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "I'm 008."}}

	for k1, _ := range xs {
		for _, v2 := range xs[k1] {
			fmt.Println(v2)
		}
	}

}
