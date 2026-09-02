package exercise

import "slices"

type Item struct {
	Name string
	Type string
}

type Player struct {
	Name      string
	Inventory []Item
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
}

func (p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if itemName == item.Name {
			p.Inventory = slices.Delete(p.Inventory, i, i+1)
			return
		}
	}
}

func (p *Player) UseItem(itemName string) {
	p.DropItem(itemName)
}
