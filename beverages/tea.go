package beverages

type Tea struct {
	price float64
}

func (h *Tea) Price() float64 {
	return h.price
}

func NewTea() *Tea {
	return &Tea{price: 1.5}
}
