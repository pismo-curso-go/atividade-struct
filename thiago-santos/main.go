// file: your-name/main.go
package main

import "fmt"

// Item represents a product in the order
type Item struct {
    Name  string
    Price float64
}

// Order represents a customer's order, with multiple items
type Order struct {
    UserID int
    Items  []Item
}

// Total returns the sum of all item prices in the order
func (o Order) Total() float64 {
    var sum float64
    for _, item := range o.Items {
        sum += item.Price
    }
    return sum
}

func main() {
    // example usage
    myOrder := Order{
        UserID: 42,
        Items: []Item{
            {"Pen", 2.50},
            {"Notebook", 15.00},
            {"Backpack", 120.00},
        },
    }

    fmt.Printf(
        "Order for user %d: total amount = $%.2f\n",
        myOrder.UserID, myOrder.Total(),
    )
}