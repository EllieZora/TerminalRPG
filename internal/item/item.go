package item

import (
	"fmt"
)

type Item struct {
	Code, Category, Name, Description string
}

func (i *Item) String() string {
	return fmt.Sprintf("%v (%v) - %v", i.Name, i.Category, i.Description)
}
