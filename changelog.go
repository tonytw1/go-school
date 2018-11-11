package main

import (
	"github.com/tonytw1/go-school/trello"
	"time"
)

type ChangelogItem struct {
	Card trello.TrelloCard
	Date time.Time
}

type CardWithActions struct {
	Card    trello.TrelloCard
	actions []trello.TrelloAction
}

func BuildChangeLog() ([]ChangelogItem, error) {
	var changelogItems []ChangelogItem

	cards, err := trello.GetCards()
	if err != nil {
		panic(err)

	} else {
		ch := make(chan CardWithActions, 100)
		for i := 0; i < len(cards); i++ {
			decorateCardWithActions(cards[i], ch)
		}

		var cardsWithActions []CardWithActions
		for i := 0; i < len(cards); i++ {
			cardsWithActions = append(cardsWithActions, <-ch)
		}

		for i := 0; i < len(cardsWithActions); i++ {
			cardWithActions := cardsWithActions[i]
			moves := cardMovesFromActions(cardWithActions.actions)
			if len(moves) > 0 {
				latestMove := moves[0]
				if latestMove.Data.ListAfter.Name == "Done" {
					item := ChangelogItem{cardWithActions.Card, latestMove.Date}
					changelogItems = append(changelogItems, item)
				}
			}
		}
	}

	return changelogItems, nil
}

func decorateCardWithActions(card trello.TrelloCard, ch chan CardWithActions) {
	actions, err := trello.GetCardActions(card.Id)
	if err != nil {
		panic(err)
	} else {
		ch <- CardWithActions{card, actions}
	}
}

func cardMovesFromActions(actions []trello.TrelloAction) []trello.TrelloAction {
	var moves []trello.TrelloAction
	for j := 0; j < len(actions); j++ {
		action := actions[j]
		if len(action.Data.ListBefore.Name) > 0 && len(action.Data.ListAfter.Name) > 0 {
			moves = append(moves, action)
		}
	}
	return moves
}
