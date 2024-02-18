package main

import (
	"fmt"
)

func main() {
	items := map[string] float64{"Paneer":16.50,"Naan":14.50}
	b := NewBill("Amar's Bill")
	b.items = items
	b.tip = 10.0
	fmt.Println(b.updateBillWithTotalAmount())
}