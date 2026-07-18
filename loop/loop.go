package main

import "fmt"

func main (){
	// for i:=0; i<5; i++{
	// 	fmt.Println(i)
	// }

	// for i:=0;i<=100; i+=10{
	// 	fmt.Println(i)
	// }

	// loop dengan if
	// for i:= 0; i<5; i++{
	// 	if i == 3{
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	for i:=0; i<5;i++{
		if i == 3 {
			break
		}

		fmt.Println(i)
	}
}
