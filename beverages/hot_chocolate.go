package beverages

type HotChocolate struct {
	price float64
}

func (h *HotChocolate) Price() float64 {
	return h.price
}

func NewHotChocolate() *HotChocolate {
	return &HotChocolate{price: 1.45}
}
