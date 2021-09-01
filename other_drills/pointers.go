package main

import "fmt"

func main() {
	fmt.Println("==================Pointers===================")
	// var a int = 10
	// fmt.Printf("Address of a variable: %x\n", &a)

	// var x int = 12345
	// var p *int
	// p = &x

	// fmt.Println("Value stored in x = ", x)
	// fmt.Println("Address of x = ", &x)
	// fmt.Println("Value stored in variable p = ", p)
	// fmt.Println("Value stored in x(*p) = ", *p)
	// fmt.Println("Address of p = ", &p)

	// fmt.Println("Value stored in x(*p) Before Changing = ", *p)

	// *p = 6789

	// fmt.Println("Value stored in x(*p) after Changing = ", x)

	i := 1
	*res := incr(&i)
	fmt.Println(*res)

}

func incr(x *int) {
	*x++
	return *x
}
