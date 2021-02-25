package beverage_machine_test

import (
	"testing"

	"github.com/myugen/go_beverages_price_refactoring_kata/beverage_machine"

	"github.com/stretchr/testify/assert"

	"github.com/myugen/go_beverages_price_refactoring_kata/beverages"
)

const precision = 0.001

func TestCoffee_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffee beverages.Beverage
	coffee = beverage_machine.NewBeverageMachine().Coffee().Make()
	assertt.InDelta(1.2, coffee.Price(), precision)
}

func TestCoffeeWithCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithCinnamon beverages.Beverage
	coffeeWithCinnamon = beverage_machine.NewBeverageMachine().Coffee().WithCinnamon().Make()
	assertt.InDelta(1.25, coffeeWithCinnamon.Price(), precision)
}

func TestCoffeeWithMilk_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilk beverages.Beverage
	coffeeWithMilk = beverage_machine.NewBeverageMachine().Coffee().WithMilk().Make()
	assertt.InDelta(1.3, coffeeWithMilk.Price(), precision)
}

func TestCoffeeWithMilkAndCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilkAndCinnamon beverages.Beverage
	coffeeWithMilkAndCinnamon = beverage_machine.NewBeverageMachine().Coffee().WithMilk().WithCinnamon().Make()
	assertt.InDelta(1.35, coffeeWithMilkAndCinnamon.Price(), precision)
}

func TestCoffeeWithMilkAndCream_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilkAndCream beverages.Beverage
	coffeeWithMilkAndCream = beverage_machine.NewBeverageMachine().Coffee().WithMilk().WithCream().Make()
	assertt.InDelta(1.45, coffeeWithMilkAndCream.Price(), precision)
}

func TestCoffeeWithMilkAndCreamAndCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilkAndCreamAndCinnamon beverages.Beverage
	coffeeWithMilkAndCreamAndCinnamon = beverage_machine.NewBeverageMachine().Coffee().WithMilk().WithCream().WithCinnamon().Make()
	assertt.InDelta(1.5, coffeeWithMilkAndCreamAndCinnamon.Price(), precision)
}

func TestTea_Price(t *testing.T) {
	assertt := assert.New(t)
	var tea beverages.Beverage
	tea = beverage_machine.NewBeverageMachine().Tea().Make()
	assertt.InDelta(1.5, tea.Price(), precision)
}

func TestTeaWithCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var teaWithCinnamon beverages.Beverage
	teaWithCinnamon = beverage_machine.NewBeverageMachine().Tea().WithCinnamon().Make()
	assertt.InDelta(1.55, teaWithCinnamon.Price(), precision)
}

func TestTeaWithMilk_Price(t *testing.T) {
	assertt := assert.New(t)
	var teaWithMilk beverages.Beverage
	teaWithMilk = beverage_machine.NewBeverageMachine().Tea().WithMilk().Make()
	assertt.InDelta(1.6, teaWithMilk.Price(), precision)
}

func TestTeaWithMilkAndCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var teaWithMilkAndCinnamon beverages.Beverage
	teaWithMilkAndCinnamon = beverage_machine.NewBeverageMachine().Tea().WithMilk().WithCinnamon().Make()
	assertt.InDelta(1.65, teaWithMilkAndCinnamon.Price(), precision)
}

func TestHotChocolate_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolate beverages.Beverage
	hotChocolate = beverage_machine.NewBeverageMachine().HotChocolate().Make()
	assertt.InDelta(1.45, hotChocolate.Price(), precision)
}

func TestHotChocolateWithCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolateWithCinnamon beverages.Beverage
	hotChocolateWithCinnamon = beverage_machine.NewBeverageMachine().HotChocolate().WithCinnamon().Make()
	assertt.InDelta(1.5, hotChocolateWithCinnamon.Price(), precision)
}

func TestHotChocolateWithCream_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolateWithCream beverages.Beverage
	hotChocolateWithCream = beverage_machine.NewBeverageMachine().HotChocolate().WithCream().Make()
	assertt.InDelta(1.6, hotChocolateWithCream.Price(), precision)
}

func TestHotChocolateWithCreamAndCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolateWithCreamAndCinnamon beverages.Beverage
	hotChocolateWithCreamAndCinnamon = beverage_machine.NewBeverageMachine().HotChocolate().WithCream().WithCinnamon().Make()
	assertt.InDelta(1.65, hotChocolateWithCreamAndCinnamon.Price(), precision)
}
