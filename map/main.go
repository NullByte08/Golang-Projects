package main

import "fmt"

func main() {
	colours := map[string]int{
		"red":   5,
		"white": 7,
	}

	for key, value := range colours {
		fmt.Println("Value for ", key, " is ", value)
	}
	fmt.Println(colours)

	//var colours map[string]string
	/*colours := make(map[string]string)
	colours["white"]="#ffffff"
	delete(colours,"white")*/
}
