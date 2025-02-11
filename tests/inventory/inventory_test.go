package inventory

import (
	"testing"

	"github.com/EllieZora/TerminalRPG/internal/inventory"
)

func TestGetNumItem(t *testing.T) {
	inv := inventory.NewInventory(1)

	resultNumItems0 := inv.GetNumItem("00")
	expectedNumItems0 := 0
	if resultNumItems0 != expectedNumItems0 {
		t.Errorf("number of items with code does not match: got %v want %v", resultNumItems0, expectedNumItems0)
	}

	inv.AddItem("00", 17)
	resultNumItems1 := inv.GetNumItem("00")
	expectedNumItems1 := 17
	if resultNumItems1 != expectedNumItems1 {
		t.Errorf("number of items with code does not match: got %v want %v", resultNumItems1, expectedNumItems1)
	}
}

func TestAddItem(t *testing.T) {
	inv := inventory.NewInventory(1)

	ok0 := inv.AddItem("00", 10)
	if !ok0 {
		t.Errorf("should allow adding items when under inventory limit")
	}

	ok1 := inv.AddItem("01", 1)
	if ok1 {
		t.Errorf("should not allow adding items when at inventory limit")
	}

	ok2 := inv.AddItem("00", 5)
	if !ok2 {
		t.Errorf("should allow adding items when item already in inventory")
	}

	ok3 := inv.AddItem("00", 0)
	if ok3 {
		t.Errorf("should not allow adding items of quantity  equal to zero")
	}

	ok4 := inv.AddItem("00", -1)
	if ok4 {
		t.Errorf("should not allow adding items of quantity less than zero")
	}

	resultNumItems := inv.GetNumItem("00")
	expectedNumItems := 15
	if resultNumItems != expectedNumItems {
		t.Errorf("number of items with code does not match: got %v want %v", resultNumItems, expectedNumItems)
	}
}

func TestRemoveItem(t *testing.T) {
	inv := inventory.NewInventory(1)

	inv.AddItem("00", 10)

	ok0 := inv.RemoveItem("00", 5)
	if !ok0 {
		t.Errorf("should allow remove items when under quantity")
	}

	resultNumItems0 := inv.GetNumItem("00")
	expectedNumItems0 := 5
	if resultNumItems0 != expectedNumItems0 {
		t.Errorf("number of items with code does not match: got %v want %v", resultNumItems0, expectedNumItems0)
	}

	ok1 := inv.RemoveItem("00", 5)
	if !ok1 {
		t.Errorf("should allow remove items to zero quantity")
	}

	resultNumItems1 := inv.GetNumItem("00")
	expectedNumItems1 := 0
	if resultNumItems1 != expectedNumItems1 {
		t.Errorf("number of items with code does not match: got %v want %v", resultNumItems1, expectedNumItems1)
	}

	ok2 := inv.RemoveItem("00", 1)
	if ok2 {
		t.Errorf("should not allow removing items when quantity is zero")
	}

	ok3 := inv.AddItem("01", 10)
	if !ok3 {
		t.Errorf("removing all items should clear inventory stack")
	}

	ok4 := inv.RemoveItem("01", 11)
	if ok4 {
		t.Errorf("should not allow removing more items than quantity")
	}

	resultNumItems2 := inv.GetNumItem("01")
	expectedNumItems2 := 10
	if resultNumItems2 != expectedNumItems2 {
		t.Errorf("number of items with code does not match: got %v want %v", resultNumItems2, expectedNumItems2)
	}
}

func TestPrint(t *testing.T) {
	inv := inventory.NewInventory(2)

	resultPrint0 := inv.Print()
	expectedPrint0 := "Your inventory is empty."
	if resultPrint0 != expectedPrint0 {
		t.Errorf("inventory print does not match: got %v want %v", resultPrint0, expectedPrint0)
	}

	inv.AddItem("00", 5)
	inv.AddItem("01", 10)

	resultPrint1 := inv.Print()
	expectedPrint1 := "You have:\n" + "1 stack(s) of item 00\n" + "1 stack(s) of item 01\n"
	if resultPrint1 != expectedPrint1 {
		t.Errorf("inventory print does not match: got %v want %v", resultPrint1, expectedPrint1)
	}
}
