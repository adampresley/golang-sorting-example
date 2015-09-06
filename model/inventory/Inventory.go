package inventory

import (
	"fmt"

	"github.com/adampresley/golang-sorting-example/model/product"
)

type Inventory struct {
	product.Product
	Price float64
}

func (inventory Inventory) String() string {
	return fmt.Sprintf("%s - %.2f", inventory.Name, inventory.Price)
}
