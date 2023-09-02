// test the functionality of dog
package dog_test

import (
	"fmt"
	"testing"

	"github.com/sagivO2150/Hands-on-exercises/old-course/errors/07/dog"
)

type dogExpect struct {
	HumanAge       int
	ExpectedDogAge int
}

// takes in a slice of type dog and cheacks to see if the Human age recived returns as the expected dog age
func TestYears(t *testing.T) {
	d := []dogExpect{
		{1, 7},
		{2, 14},
		{3, 21},
	}
	for _, v := range d {
		result := dog.Years(v.HumanAge)
		result2 := dog.YearsTwo(v.HumanAge)
		if v.ExpectedDogAge != result {
			t.Errorf("testing Years(%d), got %d, want %d", v.HumanAge, result, v.ExpectedDogAge)
		}

		if v.ExpectedDogAge != result2 {
			t.Errorf("testing Years(%d), got %d, want %d", v.HumanAge, result2, v.ExpectedDogAge)
		}

		if v.ExpectedDogAge == result2 && v.ExpectedDogAge == result {
			fmt.Println("got -", v.HumanAge, "expectes -", v.ExpectedDogAge)
		}
	}

}

// runs the dog.Years function and checks if it returns the expected dog age
func ExampleYears() {
	fmt.Println(dog.Years(2))
	// Output: 14
}

// runs the dog.YearsTwo function and checks if it returns the expected dog age
func ExampleYearsTwo() {
	fmt.Println(dog.YearsTwo(2))
	// Output: 14
}

// benchmarks the dog.Years function
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.Years(5)
	}
}

// benchmarks the dog.YearsTwo function
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.YearsTwo(5)
	}
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
