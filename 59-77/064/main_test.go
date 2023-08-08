package main

import "testing"

func TestAdd(t *testing.T) {
	want := Add(43, 16)
	got := 58
	if want != got {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", want, got)
	}

}

func TestSub(t *testing.T) {

	want := Subtract(42, 16)
	got := 26
	if want != got {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", want, got)
	}
}
