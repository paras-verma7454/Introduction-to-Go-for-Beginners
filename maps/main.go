package main

import (
	"fmt"
)

// We will be modifying this later
func CreateMap() map[string]string {
	color := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	return color
}

func AddToMap(colors map[string]string) map[string]string{
	colors["blue"] = "#0000ff"
	return colors
}

func DeleteFromMap(colors map[string]string) map[string]string{
	delete(colors, "blue")
	return colors
}

// Where we will run our code
func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(CreateMap())
	fmt.Println(AddToMap(colors))
	fmt.Println(DeleteFromMap(colors))
}
