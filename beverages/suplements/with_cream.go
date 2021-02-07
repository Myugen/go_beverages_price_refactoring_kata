package suplements

import "github.com/myugen/go_beverages_price_refactoring_kata/beverages"

type WithCream struct {
	beverage beverages.Beverage
	price    float64
}

func (b *WithCream) Price() float64 {
	return b.beverage.Price() + b.price
}

func NewWithCream(beverage beverages.Beverage) *WithCream {
	return &WithCream{
		beverage: beverage,
		price:    0.15,
	}
}
