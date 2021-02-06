package beverages_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/myugen/go_beverages_price_refactoring_kata/beverages"
)

const precision = 0.001

func TestCoffee_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffee beverages.Beverage
	coffee = beverages.NewCoffee()
	assertt.InDelta(1.2, coffee.Price(), precision)
}

func TestCoffeeWithMilk_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilk beverages.Beverage
	coffeeWithMilk = beverages.NewCoffeeWithMilk()
	assertt.InDelta(1.3, coffeeWithMilk.Price(), precision)
}

func TestCoffeeWithMilkAndCream_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilkAndCream beverages.Beverage
	coffeeWithMilkAndCream = beverages.NewCoffeeWithMilkAndCream()
	assertt.InDelta(1.45, coffeeWithMilkAndCream.Price(), precision)
}

func TestTea_Price(t *testing.T) {
	assertt := assert.New(t)
	var tea beverages.Beverage
	tea = beverages.NewTea()
	assertt.InDelta(1.5, tea.Price(), precision)
}

func TestTeaWithMilk_Price(t *testing.T) {
	assertt := assert.New(t)
	var teaWithMilk beverages.Beverage
	teaWithMilk = beverages.NewTeaWithMilk()
	assertt.InDelta(1.6, teaWithMilk.Price(), precision)
}

func TestHotChocolate_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolate beverages.Beverage
	hotChocolate = beverages.NewHotChocolate()
	assertt.InDelta(1.45, hotChocolate.Price(), precision)
}

func TestHotChocolateWithCream_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolateWithCream beverages.Beverage
	hotChocolateWithCream = beverages.NewHotChocolateWithCream()
	assertt.InDelta(1.6, hotChocolateWithCream.Price(), precision)
}
