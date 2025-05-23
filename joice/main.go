package main

import "fmt"

type Pedido struct {
	userID int
	itens  []Item
}

type Item struct {
	nome  string
	preco float64
}

func (p Pedido) valorTotal() float64 {
	valorTotal := 0.0
	for _, item := range p.itens {
		valorTotal += item.preco
	}
	return valorTotal
}

func main() {
	p := Pedido{
		userID: 1,
		itens: []Item{
			{nome: "Bandeja de morango", preco: 10.9},
			{nome: "Molho para salada", preco: 12.9},
			{nome: "Queijo meia-cura", preco: 19.0},
		},
	}

	fmt.Println("Itens do pedido:")
	for _, item := range p.itens {
		fmt.Printf("- %s -> R$ %.2f\n", item.nome, item.preco)
	}

	fmt.Printf("Valor total do pedido: R$ %.2f\n", p.valorTotal())
}
