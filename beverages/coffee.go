package beverages

type Coffee struct {
	price float64
}

func (h *Coffee) Price() float64 {
	return h.price
}

func NewCoffee() *Coffee {
	return &Coffee{price: 1.2}
}
