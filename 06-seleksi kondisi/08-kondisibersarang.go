package main

import "fmt"

func main(){
	var value = 4

	if value > 7{
		switch value{
		case 10 : 
			fmt.Println("Perfect!")
		default:
			fmt.Println("Nice!")
		}
	}else{
		if value == 5 {
			fmt.Println("not bad")
		}else if value == 3{
			fmt.Println("keep trying")
		}else{
			fmt.Println("get harder!")
			if value == 0 {
				fmt.Println("try harder")
			}
		}
	}
}