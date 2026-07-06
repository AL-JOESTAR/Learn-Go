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

	// note else di go harus satu baris sama kurung kurawal dari if kalau gak dia eror

	fmt.Println("=========== else if ===========")

	a := 14
	b := 14

	if (a < b) {
		fmt.Println("a lebih kecil dari b")
	} else if (a > b){
		fmt.Println("a lebih besar dari b")
	} else {
		fmt.Println("a dan b sama besarnya")
	}
}