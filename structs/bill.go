package main

type bill struct {
	name string
	items map[string] float64
	tip float64
	total float64
}

func NewBill(customerName string) bill {
    return bill{
		name: customerName,
		items: map[string]float64{},
		total: 0.0,
		tip: 0.0,
	}
}

func (b bill) updateBillWithTotalAmount() bill {
	for _, price := range b.items {
		b.total += price
	}
	b.total += b.tip
	return b
}