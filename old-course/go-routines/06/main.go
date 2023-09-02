package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("my os -", runtime.GOOS)
	fmt.Println("my os -", runtime.GOARCH)
}

/*
Create a program that prints out your OS and ARCH. Use the following commands to run it
● gorun
● go build
● go install
*/
