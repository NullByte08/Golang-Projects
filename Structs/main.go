package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			"alex@gmail.com",
			92341,
		},
	}
	//OR
	/*var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo = contactInfo{"alex@gmail.com", 92341}*/

	alex.updateFirstName("Alexander") // alex will be passed by value to updateFirstName function,
	// i.e. contents of alex will be copied and THEN used by updateFirstName function.
	pointerAlex := &alex
	pointerAlex.updateFirstNameUsingPointer("Alexander") // By using pointers we can work on actual values, rather than the copies.
	alex.printPerson()                                   // Actual firstName of alex will not be updated
}

func (p person) printPerson() {
	fmt.Printf("%+v", p) // %+v prints the field names also. %v prints only the values.
}

func (p person) updateFirstName(firstName string) {
	p.firstName = firstName
}

func (p *person) updateFirstNameUsingPointer(firstName string) {
	(*p).firstName = firstName
}
