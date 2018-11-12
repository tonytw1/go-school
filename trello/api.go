package trello

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const trello_api = "https://api.trello.com/1"

func GetCards(key string, token string) ([]TrelloCard, error) {
	u, e := url.Parse(trello_api + "/search")
	if e != nil {
		return nil, e

	} else {
		q := u.Query()
		q.Set("key", key)
		q.Set("token", token)
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
				var result TrelloSearchResult
				err := json.Unmarshal(content, &result)
				if err != nil {
					return nil, nil
				} else {
					return result.Cards, nil
				}
			}
		}
	}
}

func GetCardActions(id string, key string, token string) ([]TrelloAction, error) {
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
				var result []TrelloAction
				err := json.Unmarshal(content, &result)
				if err != nil {
					return nil, nil
				} else {
					return result, nil
				}
			}
		}
	}
}
