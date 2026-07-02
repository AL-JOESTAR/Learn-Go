package main

import "fmt"

func main(){
	var name string = "andi"
	var umur int = 18

	fmt.Println("nama kamu : ", name)
	fmt.Println("umur kamu : ", umur)

	// Go juga bisa langusng memprediksi apa tipe data nya

	Bulan := "september"
	Tahun := 2000

	fmt.Println("lahir bulan : ", Bulan)
	fmt.Println("pada tahun : ", Tahun)

	//konstanta
	const pi = 3.14
	const appName = "Al foundation"

	fmt.Println("nilai phi adalah : ", pi)
	fmt.Println("nama aplikasi adalah : ", appName)
}