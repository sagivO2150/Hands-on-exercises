package main

import "fmt"

func main() {

	m := make(map[string][]string)

	m[`bond_james`] = []string{`shaken, not stirred`, `martinis`, `fast cars`}
	m[`moneypenny_jenny`] = []string{`intelligence`, `literature`, `computer science`}
	m[`no_dr`] = []string{`cats`, `ice cream`, `sunsets`}

	//added another record to the map
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	//deleted a record
	delete(m, `bond_james`)

	for k, v := range m {
		fmt.Printf("the Value of %v IS - \n", k)
		fmt.Println("---------------------------------------------------")
		for k1, v1 := range v {
			fmt.Printf("'%v' and it is in spot number - %v in the string\n", v1, k1)
		}
		fmt.Println("---------------------------------------------------")

	}
}
