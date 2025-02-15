package main

import (
	"fmt"
)

// We will be modifying this later
type Person struct {
    Name string
}

type Robot struct {
    Model string
}

type Animal struct {
    Name string
}

type Speaker interface {
    Speak() string
}

func (p Person) Speak() string {
    return fmt.Sprintf("Hi! my name is %s.\n", p.Name)
}
func (p Person) Scramble(int) int {
	return 10
}

func (r Robot) Speak() string {
    return fmt.Sprintf("Hello, I am robot model %s.\n", r.Model)
}

func (a Animal) Speak() []byte {
    return []byte(a.Name)
}

func introduce(s Speaker) {
    fmt.Println(s.Speak())
}

// Where we will run our code
func main() {
    person := Person{Name: "Alice"}
    robot := Robot{Model: "RX-78"}

    personSpoke:=person.Speak()
    robotSpoke:=robot.Speak()

    fmt.Println(personSpoke)
    fmt.Println(robotSpoke)

    introduce(person)
    introduce(robot)
    var speaker Speaker
    speaker = person
    fmt.Println("Person",speaker.Speak())
    speaker = robot
    fmt.Println(speaker.Speak())
    animal := Animal{Name: "Lion"}
    fmt.Println("Animal",animal.Speak())
}
