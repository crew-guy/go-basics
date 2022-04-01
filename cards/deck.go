package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(card, i)
	}
}

func newDeck() deck {
	cardTitles := deck{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := deck{"One", "Two", "Three", "Four"}
	cards := deck{}
	for _, title := range cardTitles {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+title)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func readFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	for i, _ := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		randomNumber := r.Intn(len(d) - 1)

		d[i], d[randomNumber] = d[randomNumber], d[i]
	}
}
