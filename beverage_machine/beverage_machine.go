package beverage_machine

import (
	"github.com/myugen/go_beverages_price_refactoring_kata/beverages"
	"github.com/myugen/go_beverages_price_refactoring_kata/beverages/supplements"
)

type BeverageOption int16

const (
	BeverageOptionCoffee BeverageOption = iota + 1
	BeverageOptionTea
	BeverageOptionHotChocolate
)

type BeverageMachine struct {
	beverage beverages.Beverage
}

func (b *BeverageMachine) WithMilk() *BeverageMachine {
	b.beverage = supplements.NewWithMilk(b.beverage)
	return b
}

func (b *BeverageMachine) WithCream() *BeverageMachine {
	b.beverage = supplements.NewWithCream(b.beverage)
	return b
}

func (b *BeverageMachine) WithCinnamon() *BeverageMachine {
	b.beverage = supplements.NewWithCinnamon(b.beverage)
	return b
}

func (b *BeverageMachine) Make() beverages.Beverage {
	return b.beverage
}

func NewBeverageMachine(option BeverageOption) *BeverageMachine {
	switch option {
	case BeverageOptionCoffee:
		return &BeverageMachine{beverage: beverages.NewCoffee()}
	case BeverageOptionTea:
		return &BeverageMachine{beverage: beverages.NewTea()}
	case BeverageOptionHotChocolate:
		return &BeverageMachine{beverage: beverages.NewHotChocolate()}
	}
	panic("Option not available")
}
