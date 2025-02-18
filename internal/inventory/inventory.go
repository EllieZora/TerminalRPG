package inventory

import (
	"fmt"

	"github.com/EllieZora/TerminalRPG/internal/item"
)

type Inventory struct {
	maxSlots int
	items    map[string]*itemStack
}

type itemStack struct {
	slotsUsed, quantity int
}

func NewInventory(numSlots int) Inventory {
	return Inventory{maxSlots: numSlots, items: make(map[string]*itemStack, numSlots)}
}

func (inv *Inventory) GetNumItem(code string) int {
	stack, ok := inv.items[code]
	if !ok {
		return 0
	}

	return stack.quantity
}

func (inv *Inventory) AddItem(code string, quantity int) bool {
	if quantity <= 0 {
		return false
	}

	stack, ok := inv.items[code]
	if ok {
		// TODO - adjust slots used if quantity exceeds stack size
		stack.quantity += quantity
		return true
	} else if !inv.isFull() {
		inv.items[code] = &itemStack{slotsUsed: 1, quantity: quantity}
		return true
	}

	return false
}

func (inv *Inventory) RemoveItem(code string, quantity int) bool {
	if quantity <= 0 {
		return false
	}

	stack, ok := inv.items[code]
	if !ok || stack.slotsUsed == 0 || stack.quantity < quantity {
		return false
	}

	if stack.quantity == quantity {
		delete(inv.items, code)
	} else {
		stack.quantity -= quantity
	}

	return true
}

func (inv *Inventory) Print(store *item.Store) string {
	if len(inv.items) == 0 {
		return "Your inventory is empty."
	}

	invDescription := "You have:\n"

	for code := range inv.items {
		i, ok := store.GetItem(code)
		if !ok {
			continue
		}
		// TODO - adjust num stacks when stack size is implemented
		invDescription += fmt.Sprintf("%v stack(s) of %v\n", 1, i.String())
	}

	return invDescription
}

func (inv *Inventory) isFull() bool {
	currentSlots := 0
	for _, stack := range inv.items {
		currentSlots += stack.slotsUsed
	}

	return currentSlots >= inv.maxSlots
}
