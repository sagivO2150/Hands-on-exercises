// Package dog allows us to more fully understand dogs.
package dog

// Years converts human years to dog years.
func Years(n int) int {
	return n * 7
}

// YearsTwo converts human years to dog years.
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}

/*
Hands-on exercise #1
Start with this code. Get the code ready to BET on the code (add benchmarks, examples, tests). Run the following in this order:
● tests
● benchmarks
● coverage
● coverage shown in web browser
● examples shown in documentation in a web browser
*/
