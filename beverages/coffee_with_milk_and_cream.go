package beverages

type CoffeeWithMilkAndCream struct {
	coffee *Coffee
	price  float64
}

func (h *CoffeeWithMilkAndCream) Price() float64 {
	return h.coffee.Price() + h.price
}

func NewCoffeeWithMilkAndCream() *CoffeeWithMilkAndCream {
	return &CoffeeWithMilkAndCream{
		coffee: NewCoffee(),
		price:  0.25,
	}
}
