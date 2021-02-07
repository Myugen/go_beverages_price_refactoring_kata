package supplements

import "github.com/myugen/go_beverages_price_refactoring_kata/beverages"

type WithCinnamon struct {
	beverage beverages.Beverage
	price    float64
}

func (b *WithCinnamon) Price() float64 {
	return b.beverage.Price() + b.price
}

func NewWithCinnamon(beverage beverages.Beverage) *WithCinnamon {
	return &WithCinnamon{
		beverage: beverage,
		price:    0.05,
	}
}
