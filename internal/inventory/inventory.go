package inventory

import (
	"fmt"

	"github.com/EllieZora/TerminalRPG/internal/item"
)

type Inventory struct {
	limit  int
	stacks map[string][]int
}

func NewInventory(limit int) Inventory {
	return Inventory{limit: limit, stacks: make(map[string][]int, limit)}
}

func (inv *Inventory) GetNumItem(code string) int {
	stacks, ok := inv.stacks[code]
	if !ok {
		return 0
	}

	numItems := 0
	for _, num := range stacks {
		numItems += num
	}
	return numItems
}

func (inv *Inventory) AddItem(code string, quantity int) bool {
	if quantity <= 0 {
		return false
	}

	stacks, ok := inv.stacks[code]
	if ok {
		// TODO - add max stack quantity
		stacks[0] += quantity
		return true
	} else if !inv.isFull() {
		inv.stacks[code] = []int{quantity}
		return true
	}
	return false
}

func (inv *Inventory) RemoveItem(code string, quantity int) bool {
	if quantity <= 0 {
		return false
	}

	updatedStacks, ok := inv.stacks[code]
	if !ok || len(updatedStacks) == 0 {
		return false
	}

	remaining := quantity
	for i := len(updatedStacks) - 1; i >= 0; i-- {
		if updatedStacks[i] > remaining {
			updatedStacks[i] -= remaining
			inv.stacks[code] = updatedStacks
			return true
		} else {
			remaining -= updatedStacks[i]
			updatedStacks = updatedStacks[:len(updatedStacks)-1]
			if remaining == 0 && len(updatedStacks) == 0 {
				delete(inv.stacks, code)
				return true
			} else if remaining == 0 {
				inv.stacks[code] = updatedStacks
				return true
			} else if len(updatedStacks) == 0 {
				return false
			}
		}
	}
	return true
}

func (inv *Inventory) Print(store *item.Store) string {
	if len(inv.stacks) == 0 {
		return "Your inventory is empty."
	}

	invDescription := "You have:\n"

	for code, stack := range inv.stacks {
		i, ok := store.GetItem(code)
		if !ok {
			continue
		}
		invDescription += fmt.Sprintf("%v stack(s) of %v\n", len(stack), i.String())
	}

	return invDescription
}

func (inv *Inventory) isFull() bool {
	currentItems := 0
	for _, stacks := range inv.stacks {
		currentItems += len(stacks)
	}

	return currentItems >= inv.limit
}
