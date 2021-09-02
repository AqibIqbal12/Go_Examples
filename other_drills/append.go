package main

import "fmt"

func main() {
	sl1 := []int{234, 567, 7890, 1234, 234}
	sl2 := []int{10, 100, 1000, 10000}
	// res := append(sl1, sl2...)
	// res := append(sl1, sl2[0])
	// res := append(sl1[1:], sl2[1:]...)
	res := append(sl1[2:3], sl2[2:4]...)
	fmt.Println(res)
}
