package suplements

import "github.com/myugen/go_beverages_price_refactoring_kata/beverages"

type WithMilk struct {
	beverage beverages.Beverage
	price    float64
}

func (b *WithMilk) Price() float64 {
	return b.beverage.Price() + b.price
}

func NewWithMilk(beverage beverages.Beverage) *WithMilk {
	return &WithMilk{
		beverage: beverage,
		price:    0.1,
	}
}
