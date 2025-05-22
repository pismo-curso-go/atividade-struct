package main

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

}

func (o *Order) OrderTotal() float64 {
	var total float64
	for _, item := range o.ItemsList {
		total += item.Price
	}
	return total
}
