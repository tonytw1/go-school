package main

import (
	"awesomeProject/api"
)

func main() {
	event_types, err := api.FetchEventTypes()
	if err != nil {
		panic(err)
	} else {
		for i := 0; i < len(event_types); i++ {
			eventType := event_types[i]
			println(eventType.Name)
		}
	}
}
