package makers

import (
	"github.com/myugen/go_beverages_price_refactoring_kata/beverages"
	"github.com/myugen/go_beverages_price_refactoring_kata/beverages/supplements"
)

type CoffeeMaker struct {
	beverage beverages.Beverage
}

func (m *CoffeeMaker) WithMilk() *CoffeeMaker {
	m.beverage = supplements.NewWithMilk(m.beverage)
	return m
}

func (m *CoffeeMaker) WithCream() *CoffeeMaker {
	m.beverage = supplements.NewWithCream(m.beverage)
	return m
}

func (m *CoffeeMaker) WithCinnamon() *CoffeeMaker {
	m.beverage = supplements.NewWithCinnamon(m.beverage)
	return m
}

func (m *CoffeeMaker) Make() beverages.Beverage {
	return m.beverage
}

func NewCoffeeMaker() *CoffeeMaker {
	return &CoffeeMaker{beverage: beverages.NewCoffee()}
}
