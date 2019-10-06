package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type addItemRequest struct {
	Title string `json:"title"`
	Price int    `json:"price"`
}

type deliveryRequest struct {
	ID           int `json:"id"`
	Weight       int `json:"weight"`
	DeliveryDays int `json:"deliveryDays"`
}

func CreateHTTPServer() {
	// I'm using gorilla's router here instead of the standard http one because
	// We need to be able to parse parameters from URLs for the remove endpoint
	// Could have done this with a JSON body, but I think this API is nicer.
	router := mux.NewRouter()

	router.Handle("/add", postValidationMiddleware(http.HandlerFunc(addItemHandler)))
	router.Handle("/remove/{id:[0-9]+}", http.HandlerFunc(removeItemHandler))
	router.Handle("/clear", http.HandlerFunc(clearItemsHandler))
	router.Handle("/list", http.HandlerFunc(listItemsHandler))
	router.Handle("/delivery", postValidationMiddleware(http.HandlerFunc(deliveryHandler)))

	log.Println("Listening on 3000...")
	http.ListenAndServe(":3000", router)
}
