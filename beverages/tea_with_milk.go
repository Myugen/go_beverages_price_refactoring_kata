package beverages

type TeaWithMilk struct {
	tea   *Tea
	price float64
}

func (h *TeaWithMilk) Price() float64 {
	return h.tea.Price() + h.price
}

func NewTeaWithMilk() *TeaWithMilk {
	return &TeaWithMilk{
		tea:   NewTea(),
		price: 0.1,
	}
}
