package inventory

import (
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/EllieZora/TerminalRPG/internal/item"
)

type Inventory struct {
	Limit    int
	Contents map[uuid.UUID]item.Item
}

func (inv *Inventory) GetItem(category, name string) *item.Item {
	for _, val := range inv.Contents {
		p := &val
		if p.GetCategory() != category || p.GetName() != name {
			continue
		}
		return p
	}
	return nil
}

func (inv *Inventory) AddItem(i *item.Item) error {
	existingItem := inv.GetItem(i.GetCategory(), i.GetName())
	if existingItem != nil {
		fmt.Println(existingItem)
		existingItem.SetQuantity(existingItem.GetQuantity() + i.GetQuantity())
		fmt.Println(existingItem)
	} else if len(inv.Contents) < inv.Limit {
		newItem := i.Clone()
		inv.Contents[newItem.Id] = newItem
	} else {
		return errors.New("inventory is full")
	}

	return nil
}

func (inv *Inventory) RemoveItem(category, name string, quantity int) error {
	existingItem := inv.GetItem(category, name)
	if existingItem != nil && existingItem.GetQuantity() >= quantity {
		existingItem.SetQuantity(existingItem.GetQuantity() - quantity)
		return nil
	}

	return errors.New("not enough items in inventory")
}

func (inv Inventory) Print() string {
	if len(inv.Contents) == 0 {
		return "Your inventory is empty."
	}

	invDescription := "You have:\n"

	for _, v := range inv.Contents {
		invDescription += v.Print()
	}

	return invDescription
}
