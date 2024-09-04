package main

import "fmt"

func main() {
	var value = 2

	switch value {
	case 9:
		fmt.Println("Perfect")
	case 8:
		fmt.Println("Awesome")
	default:
		fmt.Println("Not bad")
	}

}