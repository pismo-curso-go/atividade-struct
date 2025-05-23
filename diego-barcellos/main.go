package main

import "fmt"

type Order struct {
	UserID    int64
	ItemsList []Item
}

type Item struct {
	ID    int64
	Name  string
	Price float64
}

func main() {

	order := &Order{
		UserID: 1,
		ItemsList: []Item{
			{ID: 1, Name: "Item 1", Price: 10.0},
			{ID: 2, Name: "Item 2", Price: 30.0},
		},
	}

	total := order.OrderTotal()
	fmt.Println("Total Order Amount:", total)
}

func (o *Order) OrderTotal() float64 {
	var total float64
	for _, item := range o.ItemsList {
		total += item.Price
	}
	return total
}
