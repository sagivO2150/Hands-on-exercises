package main

import "fmt"

func main() {
	fmt.Println(Para("colombia"))
}

func Para(place string) string {
	return fmt.Sprint("my idea of paradis is - ", place)
}
