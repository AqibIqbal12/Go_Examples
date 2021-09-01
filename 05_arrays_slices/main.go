package main

import "fmt"

func main() {
	// Arrays
	// var fruitArr [2]string

	// // Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Decalre and assign
	// fruitArr1 := [2]string{"Apple", "Orange"}
	// fruitArr2 := fruitArr1 //pass by value

	// fmt.Println("fruitArr1", fruitArr1)
	// fmt.Println("fruitArr2", fruitArr2)
	// fruitArr2[1] = "cherry"
	// fmt.Println("=======================")
	// fmt.Println("fruitArr1", fruitArr1)
	// fmt.Println("fruitArr2", fruitArr2)
	//  =================================================

	// with pointer
	// fruitArr1 := [2]string{"Apple", "Orange"}
	// fruitArr2 := &fruitArr1 //pass by refrence
	// // fmt.Printf("%T \n", fruitArr2)

	// fmt.Println("fruitArr1", fruitArr1)
	// fmt.Println("fruitArr2", fruitArr2)
	// fruitArr2[1] = "cherry"
	// fmt.Println("=======================")
	// fmt.Println("fruitArr1", fruitArr1)
	// fmt.Println("fruitArr2", fruitArr2)

	// fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}
	// fmt.Println(len(fruitSlice))
	// fmt.Println(fruitSlice[1:3])

	nums := []int{11, 12}
	nums = append(nums, 13, 14, 15)
	fmt.Println(nums)
}
