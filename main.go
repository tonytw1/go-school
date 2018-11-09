package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type EventType struct {
	Id string
	Name string
}

func unmarshallEventTypes(content []byte) ([]EventType, error) {
	var eventTypes []EventType
	err := json.Unmarshal(content, &eventTypes)
	return eventTypes, err
}

func fetchEventTypes() ([]EventType, error) {
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

func main() {
	event_types, err := fetchEventTypes()
	if (err != nil) {
		panic(err)
	} else {
		for i := 0; i < len(event_types); i++ {
			eventType := event_types[i]
			println(eventType.Name)
		}
	}
}
