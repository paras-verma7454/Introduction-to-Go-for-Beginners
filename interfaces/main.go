package main

import "fmt"

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
    return fmt.Sprintf("Hi, my name is %s\n", p.Name)
}
func (p Person) Scramble(int) int {
	return 10
}

func (r Robot) Speak() string {
    return fmt.Sprintf("Hi, I'm a Robot model %s\n", r.Model)
}

func (a Animal) Speak() string {
    return fmt.Sprintf("This is how animal %s speaks", a.Name)
}

func introduce(s Speaker) {
    fmt.Println(s.Speak())
}

// Where we will run our code
func main() {
    person := Person{Name: "Mike"}
    robot := Robot{Model: "T-1000"}

    // personSpoke:=person.Speak()
    // robotSpoke:=robot.Speak()

    // fmt.Println(personSpoke)
    // fmt.Println(robotSpoke)

    // introduce(person)
    // introduce(robot)
    var speaker Speaker
    speaker = person
    fmt.Println("Person",speaker.Speak())
    speaker = robot
    // fmt.Println(speaker.Speak())
    animal := Animal{Name: "Lion"}
    speakers := []Speaker{person, robot, animal}
    for _, s := range speakers {
        introduce(s)
    }
}
