package main

import (
	"fmt"
	"slices"
)

// We will be modifying this later
func CreateSlice() []string{
	slice1:= []string{"dog","cat","monkey"}
	return slice1
}

func ModifySlice(slice []string) []string{
	if len(slice) <3{
		return slice
	}
	slice[1] = "elephant"
	
}

func PopSliceValue(slice []string) []string{
	slice= slices.Delete(slice, 2,4)
	return slice

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
	
	oldSlice:=[]string{"apple", "carrot","cheese"}
	oldSlice = slices.Delete(oldSlice,0,1)
	fmt.Println(oldSlice)

	fmt.Println(CreateSlice())
	slice1:= CreateSlice()
	fmt.Println(ModifySlice(slice1))
	fmt.Println(PopSliceValue(slice1))
}
