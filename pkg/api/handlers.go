package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func internalServerError(w http.ResponseWriter, logMessage string) {
	fmt.Println(logMessage)
	http.Error(w, http.StatusText(500), 500)
}

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

	// Add the item to our store
	err = addItemToStore(request)
	if err != nil {
		internalServerError(w, "ERROR: Cannot add item to store")
	} else {
		w.Write([]byte(http.StatusText(200)))
	}
}

func removeItemHandler(w http.ResponseWriter, r *http.Request) {
	// Try and get the id from the router, and convert it to an integer
	variables := mux.Vars(r)
	id, err := strconv.Atoi(variables["id"])
	if err != nil {
		http.Error(w, "Bad Request: ID should be a number", 400)
		return
	}

	err = removeItemFromStore(id)
	if err != nil {
		internalServerError(w, "ERROR: Cannot remove item from store")
	} else {
		w.Write([]byte(http.StatusText(200)))
	}
}

func clearItemsHandler(w http.ResponseWriter, r *http.Request) {
	err := clearStore()
	if err != nil {
		internalServerError(w, "ERROR: Cannot clear store")
	} else {
		w.Write([]byte(http.StatusText(200)))
	}
}

func listItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	items, err := getItemsFromStore()
	if err != nil {
		internalServerError(w, "ERROR: Cannot get items from store")
		return
	}

	// We need to convert the items to JSON so they can be transferred over HTTP.
	bytes, err := json.Marshal(items)
	if err != nil {
		internalServerError(w, "ERROR: Cannot stringify items in store")
	} else {
		w.Write(bytes)
	}
}
