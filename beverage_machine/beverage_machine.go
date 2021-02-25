package beverage_machine

import "github.com/myugen/go_beverages_price_refactoring_kata/beverage_machine/makers"

type BeverageOption int16

type BeverageMachine struct {
	coffeeMaker       *makers.CoffeeMaker
	hotChocolateMaker *makers.HotChocolateMaker
	teaMaker          *makers.TeaMaker
}

func (b *BeverageMachine) Coffee() *makers.CoffeeMaker {
	return b.coffeeMaker
}

func (b *BeverageMachine) Tea() *makers.TeaMaker {
	return b.teaMaker
}

func (b *BeverageMachine) HotChocolate() *makers.HotChocolateMaker {
	return b.hotChocolateMaker
}

func NewBeverageMachine() *BeverageMachine {
	return &BeverageMachine{
		coffeeMaker:       makers.NewCoffeeMaker(),
		hotChocolateMaker: makers.NewHotChocolateMaker(),
		teaMaker:          makers.NewTeaMaker(),
	}
}
