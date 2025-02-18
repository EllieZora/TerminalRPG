package item

import (
	"fmt"
)

type Item struct {
	Code, Category, Name, Description                                                 string
	Slash, Bludgeon, Pierce, Fire, Ice, Lightning, Spirit                             int
	ProtSlash, ProtBludgeon, ProtPierce, ProtFire, ProtIce, ProtLightning, ProtSpirit int
}

func (i *Item) String() string {
	return fmt.Sprintf("%v (%v) - %v", i.Name, i.Category, i.Description)
}
