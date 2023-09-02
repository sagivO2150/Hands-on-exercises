package main

import "testing"

type TestType struct {
	n        int
	expected int
}

func TestFactorial(t *testing.T) {
	//test cases

	testCases := []TestType{
		{
			n:        0,
			expected: 1,
		},
		{
			n:        1,
			expected: 1,
		},
		{
			n:        5,
			expected: 120,
		},
		{
			n:        6,
			expected: 120,
		},
	}

	for _, tc := range testCases {
		result := Factorial(tc.n)
		if result != tc.expected {
			t.Errorf("factorial(%d) is -> %d, while your test is expecting - %d", tc.n, result, tc.expected)
		}
	}
}
