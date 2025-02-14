package main

import (
	"fmt"
	"slices"
)

// We will be modifying this later
func CreateSlice() []int{
	slice1:= []int{1, 2, 3, 4, 5}
	return slice1
}

func ModifySlice() []int{
	slice2:= []int{1, 2, 3, 4, 5}
	slice2= append(slice2, 6)
	return slice2
}

func PopSliceValue() []int{
	slice3:= []int{1, 2, 3, 4, 5}
	slice3 = slices.Delete(slice3, 2,4)
	return slice3
}

// Where we will run our code
func main() {
	// myColors:= []string{"red", "green", "yellow"}
	// // fmt.Println("myColors:", myColors[1])
	// myColors= append(myColors, "blue")
	// fmt.Println("myColors:", myColors[3])

	// for i, color := range myColors{
	// 	fmt.Printf("color %s at index %d\n", color, i)
	// }
	
	// oldSlice:=[]string{"apple", "carrot","cheese"}
	// oldSlice = slices.Delete(oldSlice,0,1)
	// fmt.Println(oldSlice)

	fmt.Println(CreateSlice())
	fmt.Println(ModifySlice())
	fmt.Println(PopSliceValue())
}
