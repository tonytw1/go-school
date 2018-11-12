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

func BuildChangeLog(key string, token string) ([]ChangelogItem, error) {
	var changelogItems []ChangelogItem

	cards, err := trello.GetCards(key, token)
	if err != nil {
		panic(err)

	} else {
		ch := make(chan CardWithActions)
		for i := 0; i < len(cards); i++ {
			go decorateCardWithActions(cards[i], ch, key, token)
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

func decorateCardWithActions(card trello.TrelloCard, ch chan CardWithActions, key string, token string) {
	actions, err := trello.GetCardActions(card.Id, key, token)
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
