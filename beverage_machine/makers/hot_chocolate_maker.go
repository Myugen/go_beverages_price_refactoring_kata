package makers

import (
	"github.com/myugen/go_beverages_price_refactoring_kata/beverages"
	"github.com/myugen/go_beverages_price_refactoring_kata/beverages/supplements"
)

type HotChocolateMaker struct {
	beverage beverages.Beverage
}

func (m *HotChocolateMaker) WithCream() *HotChocolateMaker {
	m.beverage = supplements.NewWithCream(m.beverage)
	return m
}

func (m *HotChocolateMaker) WithCinnamon() *HotChocolateMaker {
	m.beverage = supplements.NewWithCinnamon(m.beverage)
	return m
}

func (m *HotChocolateMaker) Make() beverages.Beverage {
	return m.beverage
}

func NewHotChocolateMaker() *HotChocolateMaker {
	return &HotChocolateMaker{beverage: beverages.NewHotChocolate()}
}
