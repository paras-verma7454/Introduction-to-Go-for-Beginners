package main

import "fmt"

type Person struct {
	Name 		string
	Age  		int
	Employed 	bool
}
// We will be modifying this later
func UpdateNameByValue(p Person, newName string) Person{
	p.Name = newName
	return p
}

func UpdateNameByReference(p *Person, name string) {
	p.Name = name
}

func MakePersonEmployed(p *Person) {
	p.Employed = true
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
