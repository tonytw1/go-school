package main

import (
	"trelloChangelog/api"
	"time"
)

type ChangelogItem struct {
	Card api.TrelloCard
	Date time.Time
}

func main() {
	changelogItems, err := buildChangeLog()
	if err != nil {
		panic(err)
	} else {
		for i := 0; i < len(changelogItems); i++ {
			item := changelogItems[i]
			println(item.Card.Name)
			println(item.Date.String())
			println(item.Card.Desc)
			println()
		}
	}
}

func buildChangeLog() ([]ChangelogItem, error) {
	var changelogItems []ChangelogItem

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
				if len(moves) > 0 {
					latestMove := moves[0]
					if latestMove.Data.ListAfter.Name == "Done" {
						item := ChangelogItem{card, latestMove.Date}
						changelogItems = append(changelogItems, item)
					}
				}
			}
		}
	}

	return changelogItems, nil
}

func cardMoves(card api.TrelloCard) ([]api.TrelloAction, error) {
	trelloActions, e := api.GetCardActions(card.Id)
	if e != nil {
		panic(e)

	} else {
		var moveActions []api.TrelloAction
		for j := 0; j < len(trelloActions); j++ {
			action := trelloActions[j]
			if len(action.Data.ListBefore.Name) > 0 && len(action.Data.ListAfter.Name) > 0 {
				moveActions = append(moveActions, action)
			}
		}
		return moveActions, nil
	}
}
