package main

import "fmt"

type Person struct {
	Name 		string
	Age  		int
	Employed 	bool
}
// We will be modifying this later
func UpdateNameByValue(Person Person, name string) Person{
	Person.Name = name
	return Person
}

func UpdateNameByReference(Person *Person, name string) {
	Person.Name = name
}

func MakePersonEmployed(person *Person) {
	person.Employed = true
}

// Where we will run our code
func main() {
	Person1 := Person{
		Name: "John",
		Age: 25,
		Employed: false,
	}
	fmt.Println(UpdateNameByValue(Person1,"Jane"))
	UpdateNameByReference(&Person1, "Jony")
	MakePersonEmployed(&Person1)

}
