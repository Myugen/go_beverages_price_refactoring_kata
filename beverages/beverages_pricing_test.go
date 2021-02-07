package beverages_test

import (
	"testing"

	"github.com/myugen/go_beverages_price_refactoring_kata/beverages/suplements"

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

func TestCoffeeWithCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithCinnamon beverages.Beverage
	coffeeWithCinnamon = suplements.NewWithCinnamon(beverages.NewCoffee())
	assertt.InDelta(1.25, coffeeWithCinnamon.Price(), precision)
}

func TestCoffeeWithMilk_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilk beverages.Beverage
	coffeeWithMilk = suplements.NewWithMilk(beverages.NewCoffee())
	assertt.InDelta(1.3, coffeeWithMilk.Price(), precision)
}

func TestCoffeeWithMilkAndCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilkAndCinammon beverages.Beverage
	coffeeWithMilkAndCinammon = suplements.NewWithCinnamon(suplements.NewWithMilk(beverages.NewCoffee()))
	assertt.InDelta(1.35, coffeeWithMilkAndCinammon.Price(), precision)
}

func TestCoffeeWithMilkAndCream_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilkAndCream beverages.Beverage
	coffeeWithMilkAndCream = suplements.NewWithCream(suplements.NewWithMilk(beverages.NewCoffee()))
	assertt.InDelta(1.45, coffeeWithMilkAndCream.Price(), precision)
}

func TestCoffeeWithMilkAndCreamAndCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var coffeeWithMilkAndCreamAndCinnamon beverages.Beverage
	coffeeWithMilkAndCreamAndCinnamon = suplements.NewWithCinnamon(suplements.NewWithCream(suplements.NewWithMilk(beverages.NewCoffee())))
	assertt.InDelta(1.5, coffeeWithMilkAndCreamAndCinnamon.Price(), precision)
}

func TestTea_Price(t *testing.T) {
	assertt := assert.New(t)
	var tea beverages.Beverage
	tea = beverages.NewTea()
	assertt.InDelta(1.5, tea.Price(), precision)
}

func TestTeaWithCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var teaWithCinnamon beverages.Beverage
	teaWithCinnamon = suplements.NewWithCinnamon(beverages.NewTea())
	assertt.InDelta(1.55, teaWithCinnamon.Price(), precision)
}

func TestTeaWithMilk_Price(t *testing.T) {
	assertt := assert.New(t)
	var teaWithMilk beverages.Beverage
	teaWithMilk = suplements.NewWithMilk(beverages.NewTea())
	assertt.InDelta(1.6, teaWithMilk.Price(), precision)
}

func TestTeaWithMilkAndCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var teaWithMilkAndCinnamon beverages.Beverage
	teaWithMilkAndCinnamon = suplements.NewWithCinnamon(suplements.NewWithMilk(beverages.NewTea()))
	assertt.InDelta(1.65, teaWithMilkAndCinnamon.Price(), precision)
}

func TestHotChocolate_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolate beverages.Beverage
	hotChocolate = beverages.NewHotChocolate()
	assertt.InDelta(1.45, hotChocolate.Price(), precision)
}

func TestHotChocolateWithCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolateWithCinnamon beverages.Beverage
	hotChocolateWithCinnamon = suplements.NewWithCinnamon(beverages.NewHotChocolate())
	assertt.InDelta(1.5, hotChocolateWithCinnamon.Price(), precision)
}

func TestHotChocolateWithCream_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolateWithCream beverages.Beverage
	hotChocolateWithCream = suplements.NewWithCream(beverages.NewHotChocolate())
	assertt.InDelta(1.6, hotChocolateWithCream.Price(), precision)
}

func TestHotChocolateWithCreamAndCinnamon_Price(t *testing.T) {
	assertt := assert.New(t)
	var hotChocolateWithCreamAndCinnamon beverages.Beverage
	hotChocolateWithCreamAndCinnamon = suplements.NewWithCinnamon(suplements.NewWithCream(beverages.NewHotChocolate()))
	assertt.InDelta(1.65, hotChocolateWithCreamAndCinnamon.Price(), precision)
}
