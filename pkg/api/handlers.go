package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func findInList(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func internalServerError(w http.ResponseWriter, logMessage string) {
	fmt.Println(logMessage)
	http.Error(w, http.StatusText(500), 500)
}

func addItemHandler(w http.ResponseWriter, r *http.Request) {
	// Try and parse this into the struct
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
		internalServerError(w, "ERROR: Cannot convert response to JSON")
	} else {
		w.Write(bytes)
	}
}

func deliveryHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Try and parse this into the struct
	decoder := json.NewDecoder(r.Body)
	var request []deliveryRequest
	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	// Ensure items are correctly formed
	ids := []int{}
	for _, v := range request {
		if v.ID == 0 || v.Weight == 0 || v.DeliveryDays == 0 {
			http.Error(w, "Bad Request: ID, Weight and DeliveryDays should be above 0", 400)
			return
		}

		if v.Weight > 9 {
			http.Error(w, "Bad Request: Weight should be under 9", 400)
			return
		}

		if findInList(ids, v.ID) {
			http.Error(w, "Bad Request: Duplicate IDs entered", 400)
			return
		}

		ids = append(ids, v.ID)
	}

	// Calculate the delivery box and return to client
	bytes, err := json.Marshal(calculateBoxes(request))
	if err != nil {
		internalServerError(w, "ERROR: Cannot convert response to JSON")
	} else {
		w.Write(bytes)
	}
}
