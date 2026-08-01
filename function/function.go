package main

import "fmt"

func myMessage(){
	fmt.Println("hai ini function")
}

// with parameter

func nama(fname string){
	fmt.Println("hello", fname , "kanedy")
}

func nameAge(fname string, age int){
	fmt.Println("hai", age, "tahun", fname, "kanedy")
}

// function return
func tambah(x int, y int) int {
	return x + y
}

// named return values

func mytambah(x int, y int) (result int) {
	result = x + y
	return
}

// story return to variable
func MyTambah(x int, y int) (result int){
	result = x + y
	return result
}

//
func UmurNama(x int, y string) (result int, txt string){
	result = x + x
	txt = y + "lebih"
	return
}

func main(){
	myMessage()
	myMessage()
	myMessage()

	nama("john")
	nama("robert")

	nameAge("robert", 18)
	nameAge("john", 20)

	fmt.Println(tambah(1,2))
	fmt.Println(mytambah(1,2))
	total := MyTambah(1,2)
	fmt.Println(total)

	fmt.Println(UmurNama(5,"tahun "))
}