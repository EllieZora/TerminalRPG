package item

import (
	"fmt"
)

type Store struct {
	cache map[string]Item
}

func NewStore(items []Item) Store {
	s := Store{cache: make(map[string]Item)}
	for _, item := range items {
		s.cache[item.Code] = item
	}
	return s
}

func (s *Store) GetItem(code string) (*Item, bool) {
	existingItem, ok := s.cache[code]
	if ok {
		return &existingItem, true
	}

	retrievedItem, ok := retrieveItem(code)
	if !ok {
		return nil, false
	}
	s.cache[code] = retrievedItem
	return &retrievedItem, true
}

// TODO - retrieve from db or return nothing
func retrieveItem(code string) (Item, bool) {
	// for now just create item
	return Item{Code: code, Category: "resource", Name: code, Description: fmt.Sprintf("Description for item %v", code)}, true
}
