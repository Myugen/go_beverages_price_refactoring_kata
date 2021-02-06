package beverages

type CoffeeWithMilk struct {
	coffee *Coffee
	price  float64
}

func (h *CoffeeWithMilk) Price() float64 {
	return h.coffee.Price() + h.price
}

func NewCoffeeWithMilk() *CoffeeWithMilk {
	return &CoffeeWithMilk{
		coffee: NewCoffee(),
		price:  0.1,
	}
}
