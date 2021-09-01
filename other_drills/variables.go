package main

import "fmt"

func main() {
	// name := "abc"
	// age := 24
	name, age := "abc", 24
	fmt.Printf("%T \n", name)
	fmt.Printf("%s \n", name)
	fmt.Printf("%T \n", age)
	fmt.Printf("%d \n", age)
}
