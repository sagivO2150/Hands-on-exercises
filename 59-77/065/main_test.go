package main

import (
	"log"
	"testing"
)

func TestPara(t *testing.T) {
	want := Para("colombia")
	get := "my idea of paradis is - colombia"

	if want != get {
		log.Fatalf("incorect value, my idea of paradis is %v and you said %v\n", want, get)
	}

}
