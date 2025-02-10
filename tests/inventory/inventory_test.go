package inventory

import (
	"testing"

	"github.com/google/uuid"

	"github.com/EllieZora/TerminalRPG/internal/inventory"
	"github.com/EllieZora/TerminalRPG/internal/item"
)

func TestGetItem(t *testing.T) {
	testCategory := "resource"
	testName := "Iron Ore"
	testQuantity := 3

	inv := inventory.Inventory{Limit: 1, Contents: make(map[uuid.UUID]item.Item)}

	resultItem0 := inv.GetItem(testCategory, testName)
	if resultItem0 != nil {
		t.Errorf("should return nil when no item with matching category and name is present in the inventory")
	}

	testItem0 := item.NewItem(testCategory, testName, "test item", testQuantity)
	inv.Contents[testItem0.Id] = testItem0

	resultItem1 := inv.GetItem(testCategory, testName)
	if resultItem1 == nil {
		t.Errorf("should find item with matching category and name")
	} else {
		if resultItem1.GetCategory() != testCategory {
			t.Errorf("should find item with matching category and name")
		}
		if resultItem1.GetName() != testName {
			t.Errorf("should find item with matching category and name")
		}
		if resultItem1.GetQuantity() != testQuantity {
			t.Errorf("should find item with matching category and name")
		}
	}
}

func TestAddItemClonesInput(t *testing.T) {
	inv := inventory.Inventory{Limit: 1, Contents: make(map[uuid.UUID]item.Item)}

	testItem0 := item.NewItem("resource", "Iron Ore", "Test iron ore.", 1)
	err0 := inv.AddItem(&testItem0)
	if err0 != nil {
		t.Errorf("should allow adding items when under inventory limit")
	}

	expectedNumItems := 1
	if len(inv.Contents) != expectedNumItems {
		t.Errorf("number of items in inventory does not match: got %v want %v", len(inv.Contents), expectedNumItems)
	}

	_, ok := inv.Contents[testItem0.Id]
	if ok {
		t.Errorf("test item was not coppied into inventory")
	}

	expectedItem0 := inv.GetItem(testItem0.GetCategory(), testItem0.GetName())
	if expectedItem0.GetCategory() != testItem0.GetCategory() {
		t.Errorf("item category does not match: got %v want %v", expectedItem0.GetCategory(), testItem0.GetCategory())
	}
	if expectedItem0.GetName() != testItem0.GetName() {
		t.Errorf("item category does not match: got %v want %v", expectedItem0.GetName(), testItem0.GetName())
	}
	if expectedItem0.GetQuantity() != testItem0.GetQuantity() {
		t.Errorf("item category does not match: got %v want %v", expectedItem0.GetQuantity(), testItem0.GetQuantity())
	}
}

func TestAddItemMultipleOfSameItem(t *testing.T) {
	inv := inventory.Inventory{Limit: 1, Contents: make(map[uuid.UUID]item.Item)}

	testItem0 := item.NewItem("resource", "Iron Ore", "Test iron ore.", 2)
	inv.AddItem(&testItem0)
	inv.AddItem(&testItem0)
	inv.AddItem(&testItem0)

	expectedItem0 := inv.GetItem(testItem0.GetCategory(), testItem0.GetName())
	if expectedItem0.GetCategory() != testItem0.GetCategory() {
		t.Errorf("item category does not match: got %v want %v", expectedItem0.GetCategory(), testItem0.GetCategory())
	}
	if expectedItem0.GetName() != testItem0.GetName() {
		t.Errorf("item category does not match: got %v want %v", expectedItem0.GetName(), testItem0.GetName())
	}
	expectedQuantity := 6
	if expectedItem0.GetQuantity() != expectedQuantity {
		t.Errorf("item category does not match: got %v want %v", expectedItem0.GetQuantity(), expectedQuantity)
	}
}

func TestAddItemErrosWhenOverLimit(t *testing.T) {
	inv := inventory.Inventory{Limit: 1, Contents: make(map[uuid.UUID]item.Item)}

	testItem0 := item.NewItem("resource", "Iron Ore", "Test iron ore.", 1)
	inv.AddItem(&testItem0)
	testItem1 := item.NewItem("resource", "Copper Ore", "Test copper ore.", 1)
	err1 := inv.AddItem(&testItem1)

	expectedError := "inventory is full"
	if err1 == nil {
		t.Errorf("test item should have overflowed inventory")
	} else if err1.Error() != expectedError {
		t.Errorf("error message does not match: got %v want %v", err1.Error(), expectedError)
	}
}

func TestPrint(t *testing.T) {
	inv := inventory.Inventory{Limit: 2, Contents: make(map[uuid.UUID]item.Item)}

	expectedEmptyPrint := "Your inventory is empty."
	if inv.Print() != expectedEmptyPrint {
		t.Errorf("inventory print did not match: got %v want %v", inv.Print(), expectedEmptyPrint)
	}

	testItem := item.NewItem("resource", "Iron Ore", "Test iron ore.", 1)
	inv.AddItem(&testItem)

	expectedPrint := "You have:\n" + testItem.Print()
	if inv.Print() != expectedPrint {
		t.Errorf("inventory print did not match: got %v want %v", inv.Print(), expectedPrint)
	}
}
