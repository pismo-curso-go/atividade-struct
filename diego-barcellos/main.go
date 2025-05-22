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
