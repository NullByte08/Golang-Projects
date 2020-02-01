package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of deck which is of type []string
type deck []string

func newDeck() deck {
	return []string{"Ace", "Clubs"}
}

func (d deck) printCardsinTheDeck() {
	for i := 0; i < len(d); i++ {
		fmt.Println(d[i])
	}
}
func deal(d deck, handSize int) deck {
	return d[:handSize]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) {
	ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteslice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteslice), ",")
	return deck(s)
}

func (d deck) shuffle() deck {

	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
