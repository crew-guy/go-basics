package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson", contactInfo: contactInfo{email: "alex@acom", zipCode: "400031"}}
	alexPointer := &alex
	alexPointer.updateName("alice")
	// Alternatively, can also call updateName just by passing the struct itself instead of the address of the struct as whenever a function requires a pointer to type as input & we just pass an instance of that type, Go will intelligently use the address of the passed instance as input & make the code work normally
	alexPointer.updateName("alakay")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
