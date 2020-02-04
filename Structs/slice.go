package main

import "fmt"

func main() {
	myslice := []string{"Hi", "There"}

	update(myslice)

	fmt.Println(myslice) //In case of slices, value is updated even if we are not using pointers
}

func update(s []string) {
	s[0] = "Hello"
}
