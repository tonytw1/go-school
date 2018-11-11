package trello

import (
	"time"
)

type TrelloSearchResult struct {
	Cards []TrelloCard
}

type TrelloCard struct {
	Id               string
	Name             string
	Desc             string
	DateLastActivity time.Time
}

type TrelloActionList struct {
	Id   string
	Name string
}

type TrelloActionData struct {
	ListBefore TrelloActionList
	ListAfter  TrelloActionList
}

type TrelloAction struct {
	Id   string
	Date time.Time
	Type string
	Data TrelloActionData
}
