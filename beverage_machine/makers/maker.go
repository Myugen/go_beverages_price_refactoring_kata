package makers

import "github.com/myugen/go_beverages_price_refactoring_kata/beverages"

type Maker interface {
	Make() beverages.Beverage
}
