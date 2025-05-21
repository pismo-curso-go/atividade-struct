package main

import "fmt"

type Order struct {
	userID int
	items  []Item
}

type Item struct {
	Name  string
	Price float64
}

func main() {
	order := Order{
		userID: 1,
		items: []Item{
			{Name: "banana", Price: 5.50},
			{Name: "burger", Price: 30.00},
			{Name: "pizza", Price: 69.99},
		},
	}

	fmt.Println("Order total value: ", order.totalValue())
}

func (o Order) totalValue() float64 {
	var total float64

	for _, item := range o.items {
		total += item.Price
	}

	return total
}
