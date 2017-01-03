package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Quote represents, well, a quote
type Quote struct {
	ID      string  `json:"id,omitempty"`
	Content string  `json:"content,omitempty"`
	Author  *Author `json:"author,omitempty"`
}

// Author is the person who said so
type Author struct {
	Name string `json:"name,omitempty"`
}

func getQuoteEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range quotes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Quote{})
}

func getQuotesEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes)
}

var quotes []Quote

func main() {
	router := mux.NewRouter()

	quotes = append(quotes, Quote{ID: "1", Content: "Don't cry because it's over, smile because it happened.", Author: &Author{Name: "Dr. Seuss"}})
	quotes = append(quotes, Quote{ID: "2", Content: "So many books, so little time.", Author: &Author{Name: "Frank Zappa"}})
	quotes = append(quotes, Quote{ID: "3", Content: "You know you're in love when you can't fall asleep because reality is finally better than your dreams.", Author: &Author{Name: "Dr. Seuss"}})

	router.HandleFunc("/quotes", getQuotesEndpoint).Methods("GET")
	router.HandleFunc("/quotes/{id}", getQuoteEndpoint).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
