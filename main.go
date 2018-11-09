package main

import (
	"awesomeProject/api"
	"strconv"
)

func main() {
	cards, err := api.GetCards()
	if err != nil {
		panic(err)
	} else {
		for i := 0; i < len(cards); i++ {
			card := cards[i]

			moves, err := cardMoves(card)
			if err != nil {
				panic(err)

			} else {
				i2 := strconv.Itoa(len(moves))
				println("Moved " + i2 + " times")
				if len(moves) > 0 {
					latestMove := moves[0]
					if latestMove.Data.ListAfter.Name == "Done" {
						println(card.Name + latestMove.Date.String())
						println(latestMove.Type + " " + " " + latestMove.Data.ListBefore.Name + " -> " + latestMove.Data.ListAfter.Name)
					}
				}
			}
			println()
		}
	}

}

func cardMoves(card api.TrelloCard) ([]api.TrelloAction, error) {

	trelloActions, e := api.GetCardActions(card.Id)
	if e != nil {
		panic(e)

	} else {

		moveActions := []api.TrelloAction{}

		for j := 0; j < len(trelloActions); j++ {
			action := trelloActions[j]
			if len(action.Data.ListBefore.Name) > 0 && len(action.Data.ListAfter.Name) > 0 {
				moveActions = append(moveActions, action)
			}
		}

		return moveActions, nil
	}

}
