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

func FetchEventTypes() ([]EventType, error) {
	resp, err := http.Get("https://api.contribly.com/1/event-types")
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
