package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// These are more integration tests and test the full lifecycle of the application
// They test, adding a valid item, adding an invalid, listing items, removing an item and clearing items.
// As the code is using an in memory database, we don't need to stub this. If it was connected to a database
// It might make sense stubbing the store.

func createGetRequest(t *testing.T, rr *httptest.ResponseRecorder, handler http.Handler, url string) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	return rr
}

func createPostRequest(t *testing.T, rr *httptest.ResponseRecorder, handler http.Handler, url string, body []byte) *httptest.ResponseRecorder {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	handler.ServeHTTP(rr, req)

	return rr
}

func TestAddItemHandlerCorrect(t *testing.T) {
	rr := createPostRequest(t, httptest.NewRecorder(), http.HandlerFunc(addItemHandler), "/add", []byte(`{"title":"Shoes","price": 10}`))

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned unexpected status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `OK`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	items, _ := getItemsFromStore()
	if len(items) == 0 {
		t.Errorf("Item failed to add to store")
	}

	// Clear store for next test
	clearStore()
}

func TestAddItemHandlerIncorrect(t *testing.T) {
	rr := createPostRequest(t, httptest.NewRecorder(), http.HandlerFunc(addItemHandler), "/add", []byte(`{}`))

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned unexpected status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestListItemHandler(t *testing.T) {
	// Add an item into the store
	_ = createPostRequest(t, httptest.NewRecorder(), http.HandlerFunc(addItemHandler), "/add", []byte(`{"title":"Shoes","price": 10}`))

	// Create a request to list the items
	rr := createGetRequest(t, httptest.NewRecorder(), http.HandlerFunc(listItemsHandler), "/list")

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned unexpected status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1,"title":"Shoes","price":10}]`
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Clear store for next test
	clearStore()
}

func TestRemoveItemHandler(t *testing.T) {
	// Add an item into the store
	_ = createPostRequest(t, httptest.NewRecorder(), http.HandlerFunc(addItemHandler), "/add", []byte(`{"title":"Shoes","price": 10}`))

	req, err := http.NewRequest("GET", "/remove/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Use a slightly different setup here because we're using mux for URL parameters
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/remove/{id}", http.HandlerFunc(removeItemHandler))
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned unexpected status code: got %v want %v",
			status, http.StatusOK)
	}

	items, err := getItemsFromStore()
	if len(items) > 0 {
		t.Errorf("Item failed to delete from store")
	}
}

func TestClearItemHandler(t *testing.T) {
	// Add a couple of things into the store
	_ = createPostRequest(t, httptest.NewRecorder(), http.HandlerFunc(addItemHandler), "/add", []byte(`{"title":"Shoes","price": 10}`))
	_ = createPostRequest(t, httptest.NewRecorder(), http.HandlerFunc(addItemHandler), "/add", []byte(`{"title":"Gloves","price": 5}`))

	// Create a request to list the items
	rr := createGetRequest(t, httptest.NewRecorder(), http.HandlerFunc(listItemsHandler), "/clear")

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned unexpected status code: got %v want %v",
			status, http.StatusOK)
	}

	items, _ := getItemsFromStore()
	if len(items) == 0 {
		t.Errorf("Store failed to clear")
	}
}
