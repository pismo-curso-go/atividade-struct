package main

import "fmt"

type Item struct {
	nome  string
	preco int
}

type Pedido struct {
	userID int
	items  []Item
}

func NovoPedido(userID int) *Pedido {
	return &Pedido{
		userID: userID,
		items:  []Item{},
	}
}

func (pedido *Pedido) AddItem(nome string, preco int) {
	item := Item{nome: nome, preco: preco}
	pedido.items = append(pedido.items, item)
}

func (pedido *Pedido) CalcularValorTotal() int {
	valorTotal := 0
	for _, item := range pedido.items {
		valorTotal += item.preco
	}

	return valorTotal
}

func main() {
	usuarioID := 1
	pedido := NovoPedido(usuarioID)
	calcularValor(*pedido)

	pedido.AddItem("teste", 10)
	calcularValor(*pedido)

	pedido.AddItem("teste2", 10)
	pedido.AddItem("teste3", 10)
	calcularValor(*pedido)
}

func calcularValor(pedido Pedido) {
	valorTotal := pedido.CalcularValorTotal()
	fmt.Printf("Valor total do pedido do usuario %d: R$%d\n", pedido.userID, valorTotal)

}
