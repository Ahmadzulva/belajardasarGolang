package main

import "fmt"

func main() {
	// versi ribet
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	// setengah ringkas
	var fourth, fifth, sixth string = "empat", "lima", "enam"

	// singkat
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	fmt.Printf("%s %s %s %s %s %s %s %s %s", first, second, third, fourth, fifth, sixth, seventh, eight, ninth)
}