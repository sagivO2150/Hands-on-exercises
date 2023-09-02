// CenteredAvg computes the average of a list of numbers
// after removing the largest and smallest values.
package mymath_test

import (
	"fmt"
	"testing"

	"github.com/sagivO2150/Hands-on-exercises/old-course/errors/09/mymath"
)

type teststruct struct {
	input  []int
	output float64
}

func TestCAvg(t *testing.T) {
	tests := []teststruct{{
		input:  []int{1, 4, 6, 8, 100},
		output: 6.0,
	},
		{
			input:  []int{0, 8, 10, 1000},
			output: 9,
		},
		{
			input:  []int{9000, 4, 10, 8, 6, 12},
			output: 9,
		},
		{
			input:  []int{123, 744, 140, 200},
			output: 170,
		},
	}
	for _, test := range tests {
		got := mymath.CenteredAvg(test.input)
		if got != test.output {
			t.Errorf("for the slice %v expected center avarage is %f, got %f", test.input, test.output, got)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(mymath.CenteredAvg([]int{123, 744, 140, 200}))
	//output: 170
}

func BenchmarkTestCAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mymath.CenteredAvg([]int{123, 744, 140, 200})
	}
}
