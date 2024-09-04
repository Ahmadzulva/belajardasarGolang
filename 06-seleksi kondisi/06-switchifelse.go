package main

import "fmt"

func main() {
	var value = 8

	switch {
	case value == 9:
		fmt.Println("Perfect")
	case (value < 9) && (value > 3):
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

}