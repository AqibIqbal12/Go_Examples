package main

import "fmt"

func main() {

	// // Creating an array
	// arr := [7]string{"This", "is", "the", "tutorial",
	// 	"of", "Go", "language"}

	// // Display array
	// fmt.Println("Array:", arr)

	// // Creating a slice
	// myslice := arr[1:6]

	// // Display slice
	// fmt.Println("Slice:", myslice)

	// // Display length of the slice
	// fmt.Printf("Length of the slice: %d", len(myslice))

	// // Display the capacity of the slice
	// fmt.Printf("\nCapacity of the slice: %d\n", cap(myslice))

	// ===================================================================

	// Creating an array
	arr := [4]string{"Geeks", "for", "Geeks", "GFG"}

	// Creating slices from the given array
	var my_slice_1 = arr[1:2]
	my_slice_2 := arr[0:]
	my_slice_3 := arr[:2]
	my_slice_4 := arr[:]

	// Display the result
	// fmt.Println("My Array: ", arr)
	fmt.Println("My Slice 1: ", my_slice_1)
	fmt.Println("My Slice 2: ", my_slice_2)
	fmt.Println("My Slice 3: ", my_slice_3)
	fmt.Println("My Slice 4: ", my_slice_4)
}
