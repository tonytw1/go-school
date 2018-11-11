package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	changelogItems, err := BuildChangeLog()
	if err != nil {
		panic(err)
	} else {
		for i := 0; i < len(changelogItems); i++ {
			item := changelogItems[i]
			fmt.Fprintf(w, item.Card.Name)
			fmt.Fprintf(w, item.Date.String())
			fmt.Fprintf(w, item.Card.Desc)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
