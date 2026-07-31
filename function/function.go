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


func main(){
	myMessage()
	myMessage()
	myMessage()

	nama("john")
	nama("robert")

	nameAge("robert", 18)
	nameAge("john", 20)

	fmt.Println(tambah(1,2))
}