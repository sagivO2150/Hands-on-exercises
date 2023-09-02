package word_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/sagivO2150/Hands-on-exercises/old-course/errors/08/word"
)

type teststruct struct {
	text     string
	numWords int
}

func TestCount(t *testing.T) {
	tests := []teststruct{
		{
			text:     "hello world",
			numWords: 2,
		},
		{
			text:     "my name is sagiv",
			numWords: 4,
		},
		{
			text:     "today i will go to the gym",
			numWords: 7,
		},
	}

	for _, test := range tests {
		got := word.Count(test.text)
		if got != test.numWords {
			t.Errorf("for the text '%v' ,Expected %v words, got %v", test.text, test.numWords, got)
		}

	}

}

func TestUseCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "hello world hello",
			expected: map[string]int{
				"hello": 2,
				"world": 1,
			},
		},
		{
			input: "apple banana apple orange banana",
			expected: map[string]int{
				"apple":  2,
				"banana": 2,
				"orange": 1,
			},
		},
	}

	for _, test := range tests {
		result := word.UseCount(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input '%s', expected %v, but got %v", test.input, test.expected, result)
		}
	}
}

func ExampleUseCount() {
	fmt.Println(word.UseCount("apple banana apple orange banana"))
	//output: map[apple:2 banana:2 orange:1]

}

func ExampleCount() {
	fmt.Println(word.Count("it is a good day"))
	// Output: 5
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		word.Count("it is a good day")
	}
}

func BenchmarkUseCount(b *testing.B) {
	input := "apple banana apple orange banana"
	for i := 0; i < b.N; i++ {
		word.UseCount(input)
	}
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
