package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Employed bool
}

// Function to create a new Person
func NewPerson(name string, age int, employed bool) Person {
	return Person{
		Name:     name,
		Age:      age,
		Employed: employed,
	}
}

func UpdateNameByValue(person Person, name string) Person {
	person.Name = name
	return person
}

func UpdateNameByReference(person *Person, name string) {
	person.Name = name
}

func MakePersonEmployed(person *Person) {
	person.Employed = true
}

// Where we will run our code
func main() {
	Person1 := NewPerson("John", 25, false)
	fmt.Println(UpdateNameByValue(Person1, "Jane"))
	UpdateNameByReference(&Person1, "Jony")
	MakePersonEmployed(&Person1)
	fmt.Println(Person1)
}
