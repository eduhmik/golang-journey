package main

import (
	"fmt"
	"strconv"
)

// define struct person
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// methods don't go within the struct

// greetings method (value receiver)

func (p Person) greetings() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// method that has a pointer receiver

func (p *Person) hasBirthday() {
	p.age++
}

// getsMarried(pointer receiver)

func (p *Person) getsMarried(spouseLastName string) {
	if p.gender == "Male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	// init person using struct
	person1 := Person{"Eduh", "Kim", "Nairobi", "Male", 24}
	person2 := Person{"Sharon", "Mundia", "Mombasa", "Female", 18}
	// print one variable

	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)
	person1.hasBirthday()
	person2.getsMarried("Michael")
	fmt.Println(person2.greetings())
}
