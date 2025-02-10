package item

import (
	"fmt"

	"github.com/google/uuid"
)

type Item struct {
	Id                          uuid.UUID
	quantity                    int
	category, name, description string
}

func NewItem(category, name, description string, quantity int) Item {
	q := quantity
	if q < 1 {
		q = 1
	}

	return Item{
		Id:          uuid.New(),
		quantity:    q,
		category:    category,
		name:        name,
		description: description,
	}
}

func (i *Item) Clone() Item {
	return NewItem(i.category, i.name, i.description, i.quantity)
}

func (i *Item) GetCategory() string {
	return i.category
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetQuantity() int {
	return i.quantity
}

func (i *Item) SetQuantity(q int) {
	i.quantity = q
}

func (i *Item) Print() string {
	return fmt.Sprintf("%v %v - %v", i.quantity, i.name, i.description)
}
