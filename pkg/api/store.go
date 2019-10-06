package api

type item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
}

// This is an in memory database, all interactions with it are abstracted behind functions here
// If we were to convert this to a persistent database, we should only need to change this interface
// The functions all return error because of this, as our interface is just an in memory db it can't actually error
// But the handlers are setup to handle this state
var items []item

func init() {
	items = []item{}
}

func findInStore(id int, data []item) int {
	for k, v := range data {
		if id == v.ID {
			return k
		}
	}

	return -1
}

func addItemToStore(request addItemRequest) error {
	id := len(items) + 1
	items = append(items, item{
		ID:    id,
		Title: request.Title,
		Price: request.Price,
	})

	return nil
}

func removeItemFromStore(id int) error {
	index := findInStore(id, items)
	if index != -1 {
		// As I'd like to keep the items ordered, in Golang we have to shift all of the elements at
		// the right of the one being deleted, by one to the left.
		items = append(items[:index], items[index+1:]...)
	}

	return nil
}

func clearStore() error {
	items = []item{}

	return nil
}

func getItemsFromStore() ([]item, error) {
	return items, nil
}
