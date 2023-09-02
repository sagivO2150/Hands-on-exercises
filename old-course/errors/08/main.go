package main

import (
	"fmt"

	"github.com/sagivO2150/Hands-on-exercises/old-course/errors/08/quote"
	"github.com/sagivO2150/Hands-on-exercises/old-course/errors/08/word"
)

func main() {
	// fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

	fmt.Println("total number of words in the string is -", word.Count(quote.SunAlso))
}

/*
Hands-on exercise #2
Start with this code. Get the code ready to BET on (add benchmarks, examples, tests) however
do not write an example for the func that returns a map; and only write a test for it as an extra challenge. Add documentation to the code. Run the following in this order:
● tests
● benchmarks
● coverage
● coverage shown in web browser
● examples shown in documentation in a web browser
*/
