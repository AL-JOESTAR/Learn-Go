package main

import "fmt"

func myMessage(){
	fmt.Println("hai ini function")
}

func nama(fname string){
	fmt.Println("hello", fname , "kanedy")
}

func nameAge(fname string, age int){
	fmt.Println("hai", age, "tahun", fname, "kanedy")
}

func main(){
	myMessage()
	myMessage()
	myMessage()

	nama("john")
	nama("robert")

	nameAge("robert", 18)
	nameAge("john", 20)
}