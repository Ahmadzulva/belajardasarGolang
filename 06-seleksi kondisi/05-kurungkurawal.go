package main

import "fmt"

func main() {
	var value = 2

	switch value {
	case 9:
		fmt.Println("Perfect")
	case 8, 7, 6, 5:
		fmt.Println("Awesome")
	default:
		{
			fmt.Println("Not bad")
			fmt.Println("You can be better")
		}
	}
}