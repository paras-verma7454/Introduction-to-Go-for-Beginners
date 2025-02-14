package main

import (
	"fmt"
)

// We will be modifying this later
func CreateSlice() []int{
	slice:= []int{1, 2, 3, 4, 5}
	return slice
}

func ModifySlice() []int{
	slice:= []int{1, 2, 3, 4, 5}
	slice= append(slice, 6)
	return slice
}

func PopSliceValue() int{
	slice:= []int{1, 2, 3, 4, 5}
	return slice[0]
}

// Where we will run our code
func main() {
	// numbersArray:= []int{}
	// var realEmptySlice []int

	// if len(numbersArray) == 0 {
	// 	fmt.Println("numbersArray is empty")
	// }
	// if len(realEmptySlice) == 0 {
	// 	fmt.Println("realEmptySlice is empty")
	// }

	// myColors:= []string{"red", "green", "yellow"}
	// // fmt.Println("myColors:", myColors[1])
	// myColors= append(myColors, "blue")
	// // fmt.Println("myColors:", myColors[3])

	// for i, color := range myColors{
	// 	fmt.Printf("color %s at index %d\n", color, i)
	// }
	fmt.Println(CreateSlice())
	fmt.Println(ModifySlice())
	fmt.Println(PopSliceValue())
}
