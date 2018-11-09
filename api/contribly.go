package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type EventType struct {
	Id   string
	Name string
}

const api_root = "https://api.contribly.com/1"

func FetchEventTypes() ([]EventType, error) {
	resp, err := http.Get(api_root + "/event-types")
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		} else {
			return unmarshallEventTypes(content)
		}
	}
}

func unmarshallEventTypes(content []byte) ([]EventType, error) {
	var eventTypes []EventType
	err := json.Unmarshal(content, &eventTypes)
	return eventTypes, err
}
