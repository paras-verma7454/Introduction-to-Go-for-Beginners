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
func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age: age,
	}
}

// Where we will run our code
func main() {
	Person1 := NewPerson("John", 25)
	fmt.Println(Person1)
	fmt.Println(UpdateNameByValue(Person1,"Jane"))
	UpdateNameByReference(&Person1, "Jony")
	MakePersonEmployed(&Person1)

}
