package makers

import (
	"github.com/myugen/go_beverages_price_refactoring_kata/beverages"
	"github.com/myugen/go_beverages_price_refactoring_kata/beverages/supplements"
)

type TeaMaker struct {
	beverage beverages.Beverage
}

func (m *TeaMaker) WithMilk() *TeaMaker {
	m.beverage = supplements.NewWithMilk(m.beverage)
	return m
}

func (m *TeaMaker) WithCinnamon() *TeaMaker {
	m.beverage = supplements.NewWithCinnamon(m.beverage)
	return m
}

func (m *TeaMaker) Make() beverages.Beverage {
	return m.beverage
}

func NewTeaMaker() *TeaMaker {
	return &TeaMaker{beverage: beverages.NewTea()}
}
