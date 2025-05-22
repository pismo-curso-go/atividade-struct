package main

import "fmt"

type Item struct {
	nome       string
	quantidade int
	preco      float64
}

type Pedido struct {
	userID int
	itens  []Item
}

func (pedido Pedido) valorTotal() float64 {
	var total float64

	for _, item := range pedido.itens {
		multiplica := item.preco * float64(item.quantidade)
		total += multiplica
	}
	return total
}

func chamarPedidos(pedido Pedido) {
	total := pedido.valorTotal()
	fmt.Printf("O valor total do pedido do usuario %d, é de %.2f\n", pedido.userID, total)
}

func main() {

	item1 := Item{nome: "Caneca", quantidade: 2, preco: 10.99}
	item2 := Item{nome: "Cama", quantidade: 1, preco: 210.50}
	item3 := Item{nome: "Mouse", quantidade: 3, preco: 30.20}

	pedido1 := Pedido{userID: 1, itens: []Item{item1, item3}}
	pedido2 := Pedido{userID: 2, itens: []Item{item2}}

	pedidos := []Pedido{pedido1, pedido2}

	for _, pedido := range pedidos {
		chamarPedidos(pedido)
	}
}
