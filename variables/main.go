package main

import "fmt"

// We will be modifying this later
func VariableDeclaration() string {
	return "Hello, World!"
}

func StringConcatenation(firstArg string, secondArg string) string {
	return firstArg + secondArg
}

func IncrementInt(count int) int{
	return count+1
}

// Where we will run our code
func main() {
	// var msg string
	// var count int
	// var isActive bool = true

	// fmt.Println("msg:", msg)
	// fmt.Println("count:", count)
	// fmt.Println("isActive:", isActive)

	// myInt := 10
	// myFloat := 2.5

	// sum:= myInt + int(myFloat)
	// fmt.Println("Sum:",sum)

    // myString:= "this is my string"
	// myStringByte:= []byte(myString)
	// fmt.Println("my String in Bytes:",myStringByte)

	fmt.Println(VariableDeclaration())
	fmt.Println(StringConcatenation("Hello", "World"))
	fmt.Println(IncrementInt(10))
}
