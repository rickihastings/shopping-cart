package api

import (
	"encoding/json"
	"net/http"
)

func addItemHandler(w http.ResponseWriter, r *http.Request) {
	// Try and parse this into the struct, this code could be reusable but golang doesn't have generics
	decoder := json.NewDecoder(r.Body)
	var request addItemRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	// Ensure title is not empty
	if request.Title == "" {
		http.Error(w, "Bad Request: Missing Title", 400)
		return
	}

	w.Write([]byte(http.StatusText(200)))
}

func removeItemHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func clearItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
