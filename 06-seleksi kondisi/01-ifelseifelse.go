package main

import "fmt"

func main(){
	var value = 11

	if value == 10 {
		fmt.Println("nilai sempurna")
	}else if value > 7 {
		fmt.Println("nilai baik")
	}else if value > 5 {
		fmt.Println("nilai rata rata")
	}else{
		fmt.Println("nilai kurang")
	}
}