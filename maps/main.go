package main

import (
	"fmt"
)

// We will be modifying this later
func CreateMap(key1, key2, val1,val2 string) map[string]string{
	return map[string]string{
		key1: val1,
		key2: val2,
	}
}

func AddToMap(color map[string]int, newKey string,newValue int ) map[string]int{
	color[newKey] = newValue
	return color
}

func DeleteFromMap(color map[string]string, key string) map[string]string{
	delete(color,key)
	return color
}

// Where we will run our code
func main() {
	fmt.Println(CreateMap("First", "Last", "Jon", "Doe"))
	fmt.Println(AddToMap(map[string]int{"red": 1, "blue": 2, "green": 3},"yellow",4))
	fmt.Println(DeleteFromMap(map[string]string{"red": "1", "blue": "2", "green": "3"},"blue"))
}
