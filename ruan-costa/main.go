package main

import "fmt"

type Pedido struct {
	userID   int64
	itemList []Item
}

type Item struct {
	name  string
	price float64
}

func main() {

	pedido1 := Pedido{
		userID: 1,
		itemList: []Item{
			{name: "Lápis", price: 2.50},
			{name: "Borracha", price: 5.00},
			{name: "Caderno", price: 10.00},
		},
	}

	fmt.Println(pedido1.valorTotal())
}

func (p Pedido) valorTotal() float64 {
	var valorTotal float64

	for _, item := range p.itemList {
		valorTotal += item.price
	}

	return valorTotal
}
