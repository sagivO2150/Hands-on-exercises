// takes a string and counts the number of time a each word appears and the total number of words in the string
package word

import (
	"strings"
)

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// returns the number of words in total for the entire string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
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
