package main

import (
	"awesomeProject/api"
)

func main() {
	cards, err := api.GetCards()
	if err != nil {
		panic(err)
	} else {
		for i := 0; i < len(cards); i++ {
			println(cards[i].Name)
			println(cards[i].DateLastActivity.String())
		}
	}
}
