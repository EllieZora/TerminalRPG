package item

import (
	"fmt"
)

type Store struct {
	cache map[string]*Item
}

func NewStore(items []*Item) Store {
	s := Store{cache: make(map[string]*Item)}
	for _, item := range items {
		s.cache[item.Code] = item
	}
	return s
}

// TODO - remove when db implemented
func NewStoreDefault(items []Item) Store {
	s := Store{cache: setupDefaultItems()}
	for _, item := range items {
		s.cache[item.Code] = &item
	}
	return s
}

func (s *Store) GetItem(code string) (*Item, bool) {
	existingItem, ok := s.cache[code]
	if ok {
		return existingItem, true
	}

	retrievedItem, ok := retrieveItem(code)
	if !ok {
		return nil, false
	}
	s.cache[code] = &retrievedItem
	return &retrievedItem, true
}

// TODO - retrieve from db or return nothing
func retrieveItem(code string) (Item, bool) {
	// for now just create item
	return Item{Code: code, Category: "resource", Name: code, Description: fmt.Sprintf("Description for item %v", code)}, true
}

// TODO - setup db with default items
func setupDefaultItems() map[string]*Item {
	defaultItems := make(map[string]*Item)

	createAndAddItem(defaultItems, "0000", "ore", "Copper Ore", "A lump of copper ore.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0001", "ore", "Tin Ore", "A lump of tin ore.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0002", "ore", "Iron Ore", "A lump of iron ore.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0003", "ore", "Mythril Ore", "A lump of mythril ore.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0004", "ore", "Adamant Ore", "A lump of adamant ore.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0005", "ore", "Orichalcum Ore", "A lump of orichalcum ore.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0006", "ore", "Silver Ore", "A lump of silver ore.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0007", "ore", "Gold Ore", "A lump of gold ore.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0008", "ore", "Platinum Ore", "A lump of platinum ore.", 0, 0, 0, 0, 0, 0)

	createAndAddItem(defaultItems, "0009", "bar", "Bronze Bar", "A bar of bronze metal.", 15, 10, 10, 10, 10, 5)
	createAndAddItem(defaultItems, "0010", "bar", "Iron Bar", "A bar of iron metal.", 12, 11, 17, 15, 15, 15)
	createAndAddItem(defaultItems, "0011", "bar", "Mythril Bar", "A bar of mythril metal.", 10, 14, 10, 20, 25, 20)
	createAndAddItem(defaultItems, "0012", "bar", "Adamant Bar", "A bar of adamant metal.", 19, 19, 15, 18, 20, 20)
	createAndAddItem(defaultItems, "0013", "bar", "Orichalcum Bar", "A bar of orichalcum metal.", 23, 21, 23, 25, 25, 25)
	createAndAddItem(defaultItems, "0014", "bar", "Silver Bar", "A bar of silver metal.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0015", "bar", "Gold Bar", "A bar of gold metal.", 0, 0, 0, 0, 0, 0)
	createAndAddItem(defaultItems, "0016", "bar", "Platinum Bar", "A bar of platinum metal.", 0, 0, 0, 0, 0, 0)

	return defaultItems
}

func createAndAddItem(items map[string]*Item, code, category, name, description string, slash, bludgeon, pierce, protSlash, protBludgeon, ProtPierce int) {
	items[code] = &Item{
		Code:         code,
		Category:     category,
		Name:         name,
		Description:  description,
		Slash:        slash,
		Bludgeon:     bludgeon,
		Pierce:       pierce,
		ProtSlash:    protSlash,
		ProtBludgeon: protBludgeon,
		ProtPierce:   ProtPierce,
	}
}
