package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type EventType struct {
	Id   string
	Name string
}

type TrelloSearchResult struct {
	Cards []TrelloCard
}

type TrelloCard struct {
	Id   string
	Name string
	Desc string
}

const trello_api = "https://api.trello.com/1"

func GetCards() ([]TrelloCard, error) {
	u, e := url.Parse(trello_api + "/search")
	if e != nil {
		return nil, e

	} else {
		q := u.Query()
		q.Set("key", "")
		q.Set("token", "")
		q.Set("query", "label:changelog")
		q.Set("cards_limit", "10")
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

func unmarshallTrelloSearch(content []byte) (TrelloSearchResult, error) {
	var result TrelloSearchResult
	err := json.Unmarshal(content, &result)
	return result, err
}
