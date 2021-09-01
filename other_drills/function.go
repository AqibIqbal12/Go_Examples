package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func main() {

	var num float64 = 2
	result, err := sqrt(num)
	if err == nil {
		// fmt.Printf("square root of num %f\n", result)
		fmt.Println("square root of "+strconv.FormatFloat(num, 'f', 0, 32)+" is:", result)
	} else {
		fmt.Println(err)
	}

}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Negative number!!!!!!")
	}
	return math.Sqrt(x), nil
}
