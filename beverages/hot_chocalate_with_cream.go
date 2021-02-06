package beverages

type HotChocolateWithCream struct {
	hotChocolate *HotChocolate
	price        float64
}

func (h *HotChocolateWithCream) Price() float64 {
	return h.hotChocolate.Price() + h.price
}

func NewHotChocolateWithCream() *HotChocolateWithCream {
	return &HotChocolateWithCream{
		hotChocolate: NewHotChocolate(),
		price:        0.15,
	}
}
