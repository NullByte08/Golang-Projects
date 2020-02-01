
package main

import (
	"os"
	"testing"
)


func TestNewDeck(t *testing.T){

	d:=[]string{"Ace","Hearts"}

	if len(d)!= 3{
		t.Errorf("Expected 3 but got %v",len(d))
	}
}

func TestingSaveToFileAndNewDeckFromFile(t *testing.T){
	os.Remove("_decktesting")

	deck:=newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck:=newDeckFromFile("_decktesting")

	if len(loadedDeck)!=2 {
		t.Errorf("Expected 2 , got %v",len(loadedDeck))
	}

	os.Remove("_decktesting")
}