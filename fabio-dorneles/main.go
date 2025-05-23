package main

import "fmt"

type CadaItemDaLista struct {
	Nome  string
	Preco float64
}

type Pedido struct {
	UserID string
	Item   []CadaItemDaLista
}

/*
Esta função ValorTotal() recebe um slice de itens e retorna o valor total.
Ela percorre todos os itens do slice e soma os preços para calcular o valor total do pedido.
O tipo de dado do slice é CadaItemDaLista, que contém o preço de cada item.
*/
func ValorTotal(item []CadaItemDaLista) float64 {
	total := 0.0
	for _, i := range item {
		total += i.Preco
	}
	return total
}

func main() {
	pedido1 := Pedido{
		UserID: "001",
		Item: []CadaItemDaLista{
			{Nome: "Produto 1", Preco: 10.50},
			{Nome: "Produto 2", Preco: 20.75},
			{Nome: "Produto 3", Preco: 5.99},
		},
	}

	pedido2 := Pedido{
		UserID: "002",
		Item: []CadaItemDaLista{
			{Nome: "Produto 4", Preco: 15.00},
			{Nome: "Produto 5", Preco: 30.00},
			{Nome: "Produto 6", Preco: 12.50},
		},
	}

	fmt.Printf("Valor total do pedido do usuário %s: R$ %.2f\n", pedido1.UserID, ValorTotal(pedido1.Item))
	fmt.Printf("Valor total do pedido do usuário %s: R$ %.2f\n", pedido2.UserID, ValorTotal(pedido2.Item))
}
