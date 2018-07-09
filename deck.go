package main

// no comma
import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	// this is a slice, not array
	cardSuites := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// unused vars replaced with _
	for _, cs := range cardSuites {
		for _, cv := range cardValues {
			cards = append(cards, cv+" of "+cs)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Go allows multiple return values
func deal(d deck, handSize int) (deck, deck) {
	// slicing slices
	return d[:handSize], d[handSize:]
}

func (d deck) deckToString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// if nothing went wrong, error will be 'nil'
	bs, error := ioutil.ReadFile(filename)

	if error != nil {
		// prints out error from ioutil
		fmt.Println("Error: ", error)

		// exit with something error happening
		os.Exit(1)
	}

	strSlice := strings.Split(string(bs), ",")

	return deck(strSlice)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}
