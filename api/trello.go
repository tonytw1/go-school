package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type EventType struct {
	Id   string
	Name string
}

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

const trello_api = "https://api.trello.com/1"
const key = ""
const token = ""

func GetCards() ([]TrelloCard, error) {
	u, e := url.Parse(trello_api + "/search")
	if e != nil {
		return nil, e

	} else {
		q := u.Query()
		q.Set("key", key)
		q.Set("token", token)
		q.Set("query", "label:changelog")
		q.Set("cards_limit", "3")
		q.Set("modelTypes", "cards")
		u.RawQuery = q.Encode()
		resp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		} else {
			defer resp.Body.Close()
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			} else {
				result, err := unmarshallTrelloSearch(content)
				if err != nil {
					return nil, nil
				} else {
					return result.Cards, nil
				}
			}
		}
	}
}

func GetCardActions(id string) ([]TrelloAction, error) {
	u, e := url.Parse(trello_api + "/cards/" + id + "/actions")
	if e != nil {
		return nil, e
	} else {
		q := u.Query()
		q.Set("key", key)
		q.Set("token", token)
		u.RawQuery = q.Encode()
		resp, err := http.Get(u.String())
		if err != nil {
			return nil, err
		} else {
			defer resp.Body.Close()
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			} else {
				result, err := unmarshallTrelloActions(content)
				if err != nil {
					return nil, nil
				} else {
					return result, nil
				}
			}
		}
	}
}

func unmarshallTrelloSearch(content []byte) (TrelloSearchResult, error) {
	var result TrelloSearchResult
	err := json.Unmarshal(content, &result)
	return result, err
}

func unmarshallTrelloActions(content []byte) ([]TrelloAction, error) {
	var result []TrelloAction
	err := json.Unmarshal(content, &result)
	return result, err
}
