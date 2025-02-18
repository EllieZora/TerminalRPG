package item

import (
	"testing"

	"github.com/EllieZora/TerminalRPG/internal/item"
)

func TestGetItem(t *testing.T) {
	item0 := item.Item{Code: "00", Category: "Cat 0", Name: "Item 0", Description: "test 0"}
	item1 := item.Item{Code: "01", Category: "Cat 1", Name: "Item 1", Description: "test 1"}
	item2 := item.Item{Code: "02", Category: "Cat 2", Name: "Item 2", Description: "test 2"}
	store := item.NewStore([]*item.Item{&item0, &item1, &item2})

	existingItem0, ok := store.GetItem("00")
	if !ok {
		t.Errorf("Item with code 00 should be retrievable")
	} else {
		if existingItem0.Code != "00" {
			t.Errorf("code does not match: got %v want %v", existingItem0.Code, "00")
		}
	}

	existingItem1, ok := store.GetItem("03")
	if !ok {
		t.Errorf("Item with code 03 should be retrievable even if not in cache")
	} else {
		if existingItem1.Code != "03" {
			t.Errorf("code does not match: got %v want %v", existingItem1.Code, "03")
		}
	}
}
