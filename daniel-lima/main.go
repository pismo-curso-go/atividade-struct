package main

import "fmt"

type Item struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Pedido struct {
	UserID string
	Itens  []Item
}

func (p Pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.Itens {
		total += item.Preco * float64(item.Quantidade)
	}
	return total
}

func main() {
	// Itens para usar nos testes
	itens := map[string]Item{
		"actionFigure": {"Action Figure", 925.99, 2},
		"funkoPop":     {"Funko Pop One Piece", 79.50, 3},
		"volante":      {"Logitech G29", 1895.90, 1},
	}

	// cenarios de teste usando mapa com nome do cenarios como chave
	cenarios := map[string]struct {
		pedido   Pedido
		esperado float64
	}{
		"Pedido com 2 Action Figures e 3 Funko Pops": { //
			pedido: Pedido{
				UserID: "Cliente01",
				Itens:  []Item{itens["actionFigure"], itens["funkoPop"]},
			},
			esperado: (925.99 * 2) + (79.50 * 3),
		},
		"Pedido com 1 Volante e 2 Action Figures": {
			pedido: Pedido{
				UserID: "Cliente02",
				Itens:  []Item{itens["volante"], itens["actionFigure"]},
			},
			esperado: (1895.90 * 1) + (925.99 * 2),
		},
		"Pedido vazio": {
			pedido: Pedido{
				UserID: "Cliente03",
				Itens:  []Item{},
			},
			esperado: 0.0,
		},
		"Pedido com um único item": {
			pedido: Pedido{
				UserID: "Cliente04",
				Itens:  []Item{itens["funkoPop"]},
			},
			esperado: 79.50 * 3,
		},
	}

	// Executando os testes
	fmt.Println("🚀 Iniciando testes de pedidos...")
	fmt.Println("----------------------------------")

	passaram := 0
	totalTestes := len(cenarios)

	for nomeTeste, tc := range cenarios {
		resultado := tc.pedido.valorTotal()
		if resultado == tc.esperado {
			fmt.Printf("✅ %s - PASSOU\n", nomeTeste)
			fmt.Printf("   👤 UserID: %s\n", tc.pedido.UserID)
			fmt.Printf("   💰 Esperado: R$ %.2f | Obtido: R$ %.2f\n\n", tc.esperado, resultado)
			passaram++
		} else {
			fmt.Printf("❌ %s - FALHOU\n", nomeTeste)
			fmt.Printf("   👤 UserID: %s\n", tc.pedido.UserID)
			fmt.Printf("   💰 Esperado: R$ %.2f | Obtido: R$ %.2f\n\n", tc.esperado, resultado)
		}
	}

	fmt.Println("----------------------------------")
	fmt.Printf("📊 Resultado: %d/%d testes passaram\n", passaram, totalTestes)
	if passaram == totalTestes {
		fmt.Println("🎉 Todos os testes passaram com sucesso!")
	} else {
		fmt.Printf("⚠️  %d testes falharam\n", totalTestes-passaram)
	}
}
