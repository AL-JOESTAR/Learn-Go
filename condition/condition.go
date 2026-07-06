package main

import "fmt"

func main() {
	if 20 > 18{
		fmt.Println("20 lebih besar dari 18")
	}

	x := 10
	y := 20
	if x < y {
		fmt.Println("x lebih kecil dari y")
	}

	if y > x {
		fmt.Println("y lebih besar daripada x")
	}

	fmt.Println("=========== if else ===========")

	time := 20

	if (time < 18) {
		fmt.Println("selamat siang")
	} else {
		fmt.Println("selmat malam tuan")
	}
}