package shopitem

import (
	"fmt"

	"github.com/hyson007/goschool/assignment1/category"
)

type ShopItem struct {
	Category int     `json:"category"`
	Quantity int     `json:"quantity"`
	Cost     float64 `json:"cost"`
	Name     string  `json:"name"`
}

func (s ShopItem) String() string {
	return fmt.Sprintf("%s - {%d %d %d}\n", s.Name, s.Category, s.Quantity, int(s.Cost))
}

func (s ShopItem) PrintItem() string {
	return fmt.Sprintf("Catgory: %s - Item: % s Quantity: %d Unit Cost: %.1f\n",
		category.CategorySlice[s.Category], s.Name, s.Quantity, s.Cost)
}

func (s *ShopItem) ModifyItem(category, quantity, idx int, cost float64, name string) ShopItem {

	s.Name = name
	s.Category = category
	s.Cost = cost
	s.Quantity = quantity

	fmt.Println("hit")
	fmt.Println("name", name)
	fmt.Println(s)
	return *s
}

func (s *ShopItem) UpdateCategory() {
	s.Category -= 1
}

func NewShopItem(category, quantity int, cost float64, name string) ShopItem {
	return ShopItem{Category: category, Quantity: quantity, Cost: cost, Name: name}
}
