package main

import"fmt"

func main(){
	const(
		square = "kotak"
		isToday bool = true
		numeric uint8 = 1
		floatNum = 2.2
	)
	fmt.Println(square)
	fmt.Println(isToday)
	fmt.Println(numeric)
	fmt.Println(floatNum)
	fmt.Println("============")

	const(
		a = "konstanta"
		b
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("============")

	const(
		today = "selasa"
		sekarang
		isToday2 = true
	)
	fmt.Println(today)
	fmt.Println(sekarang)
	fmt.Println(isToday2)
	fmt.Println("============")

	const satu, dua = 1, 2
	const tiga, empat string = "3", "4"

	fmt.Println(satu,dua)
	fmt.Println(tiga,empat)
}