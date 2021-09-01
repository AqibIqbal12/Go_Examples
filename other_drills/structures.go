package main

import "fmt"

type Employee struct {
	name  string
	email string
	age   int
}

func main() {
	// e1 := Employee{name: "abc", email: "abc", age: 12}

	var Employee1 Employee
	Employee1.name = "ali"
	Employee1.email = "ali@gmail.com"
	Employee1.age = 20

	fmt.Printf("Employee1 name: %s\n", Employee1.name)
	fmt.Printf("Employee1 email: %s\n", Employee1.email)
	fmt.Printf("Employee1 age: %d\n", Employee1.age)

	println("==========================================")

	var Employee2 Employee
	Employee2.name = "hamza"
	Employee2.email = "hamza@gmail.com"
	Employee2.age = 22

	fmt.Printf("Employee2 name: %s\n", Employee2.name)
	fmt.Printf("Employee2 email: %s\n", Employee2.email)
	fmt.Printf("Employee2 age: %d\n", Employee2.age)
}
