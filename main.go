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

func unmarshallEventTypes(content []byte) []EventType {
	var eventTypes []EventType
	err := json.Unmarshal(content, &eventTypes)
	if err != nil {
		panic(err)
	} else {
		return eventTypes
	}
}

func fetchEventTypes() []EventType {
	resp, err := http.Get("https://api.contribly.com/1/event-types")
	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		} else {
			return unmarshallEventTypes(content)
		}
	}
}

func main() {
	event_types := fetchEventTypes()
	for i := 0; i < len(event_types); i++ {
		eventType := event_types[i]
		println(eventType.Name)
	}
}
