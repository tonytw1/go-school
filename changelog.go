package main

import (
	"time"
	"trelloChangelog/trello"
)

type ChangelogItem struct {
	Card trello.TrelloCard
	Date time.Time
}

func BuildChangeLog() ([]ChangelogItem, error) {
	var changelogItems []ChangelogItem

	cards, err := trello.GetCards()
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

func cardMoves(card trello.TrelloCard) ([]trello.TrelloAction, error) {
	trelloActions, e := trello.GetCardActions(card.Id)
	if e != nil {
		panic(e)

	} else {
		var moveActions []trello.TrelloAction
		for j := 0; j < len(trelloActions); j++ {
			action := trelloActions[j]
			if len(action.Data.ListBefore.Name) > 0 && len(action.Data.ListAfter.Name) > 0 {
				moveActions = append(moveActions, action)
			}
		}
		return moveActions, nil
	}
}
