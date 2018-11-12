package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	key := os.Getenv("TRELLO_KEY")
	token := os.Getenv("TRELLO_TOKEN")

	changelogItems, err := BuildChangeLog(key, token)
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
