package main

import "fmt"

func myMessage(){
	fmt.Println("hai ini function")
}

func nama(fname string){
	fmt.Println("hello", fname , "kanedy")
}

func main(){
	myMessage()
	myMessage()
	myMessage()

	nama("john")
	nama("robert")
}