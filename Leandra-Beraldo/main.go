package main

import "fmt"

// Definindo a struct Item
type Item struct {
	Nome  string
	Preco float64
}

// Definindo a struct Pedido
type Pedido struct {
	UserID int
	Itens  []Item
}

func (pedido Pedido) valorTotal() float64 {
	var total float64
	for _, item := range pedido.Itens {
		total += item.Preco
	}
	return total
}

func main() {
	pedido := Pedido{
		UserID: 1,
		Itens: []Item{
			{"Melancia", 13.00},
			{"Lichia", 6.00},
			{"Uva", 13.00},
			{"Manga", 3.50},
		},
	}

	fmt.Printf("Valor total do pedido: R$ %.2f", pedido.valorTotal())
}
