package main

import (
	"encoding/csv"
	"fmt"
	"math/rand/v2"
	"os"
)

type deck []string

func newDeck() (d deck) {
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			d = append(d, value+" of "+suit)
		}
	}

	return
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating:", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writerErr := writer.Write(d)
	if writerErr != nil {
		fmt.Println("Error writing:", writerErr)
		os.Exit(1)
	}

	writer.Flush()
}

func newDeckFromFile(filename string) deck {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening:", err)
		os.Exit(1)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading:", err)
		os.Exit(1)
	}

	return records[0]
}

func shuffle(d deck) deck {
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})

	return d
}
