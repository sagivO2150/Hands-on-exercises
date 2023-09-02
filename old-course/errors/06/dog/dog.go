// dog package is a package that takes a number of humen years and returns the number of dog years
package Dog

var sagivOron int

// Years func takes n int as the number of human years and returns the number of dog years
func YEARS(n int) int {
	return n * 7
}

/*
Hands-on exercise #1
Create a dog package. The dog package should have an exported func “Years” which takes
human years and turns them into dog years (1 human year = 7 dog years). Document your code with comments. Use this code in func main.
● run your program and make sure it works
● run a local server with godoc and look at your documentation.
*/
