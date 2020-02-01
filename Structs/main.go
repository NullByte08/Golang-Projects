package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {

	//alex := person{firstName: "Alex", lastName: "Anderson"}

	//OR

	var alex person
	alex.firstName="Alex"
	alex.lastName="Anderson"


	fmt.Println(alex)
}
