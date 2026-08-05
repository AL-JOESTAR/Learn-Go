package main

import "fmt"

type person struct{
	name string
	age int
	job string
	salary int	
}

func main () {
	var pekerja person
	var pekerja2 person

	pekerja.name = "agus"
	pekerja2.name = "supri"

	pekerja.age = 24
	pekerja2.age = 21

	pekerja.job = "programmer"
	pekerja2.job = "progamer"

	pekerja.salary = 10000000
	pekerja2.salary = 4000000

	fmt.Println("name", pekerja.name)
	fmt.Println("age", pekerja.age)
	fmt.Println("job", pekerja.job)
	fmt.Println("salary", pekerja.salary)

	fmt.Println("=========================")

	fmt.Println("name", pekerja2.name)
	fmt.Println("age", pekerja2.age)
	fmt.Println("job", pekerja2.job)
	fmt.Println("salary", pekerja2.salary)
}