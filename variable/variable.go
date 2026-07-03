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

	// Boolean
	var b1 bool = true
	var b2 = true
	var b3 bool
	b4 := true


	fmt.Println("boolean is : ", b1)
	fmt.Println("boolean is : ", b2)
	fmt.Println("boolean is : ", b3)
	fmt.Println("boolean is : ", b4)
	
	fmt.Println("===== Float 32 =======")
	var x float32 = 123.78
	var y float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v", y, y)

	fmt.Println("===== Float 64 =======")
	var x64 float64 = 1.7e+308
  	fmt.Printf("Type: %T, value: %v", x64, x64)

}