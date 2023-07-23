package main

import (
	"fmt"

	"github.com/sagivO2150/puppy"
)

func main() {
	fmt.Println("Hello world")
	fmt.Printf("a puppy goes %v \n", puppy.Barks())
	fmt.Printf("but, %v \n", puppy.BigBarks())
}
