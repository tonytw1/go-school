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
			card := cards[i]
			println(card.Name)
			trelloActions, e := api.GetCardActions(card.Id)
			if e != nil {
				panic(e)

			} else {
				for j := 0; j < len(trelloActions); j++ {
					action := trelloActions[j]
					println(action.Type + " " + action.Date.String() + " " + action.Data.ListBefore.Name + " -> " + action.Data.ListAfter.Name)
				}
			}
			println()
		}
	}

}
