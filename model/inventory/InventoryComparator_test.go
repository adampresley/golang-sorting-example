package inventory

import (
	"sort"
	"testing"

	"github.com/adampresley/golang-sorting-example/model/product"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInventoryComparator(t *testing.T) {
	Convey("Inventory comparator", t, func() {
		Convey("Sorts by name", func() {
			actual := []Inventory{
				{Product: product.Product{Name: "Widget 2", Size: product.NA}, Price: 59.75},
				{Product: product.Product{Name: "Widget 1", Size: product.NA}, Price: 199.99},
				{Product: product.Product{Name: "Dr. Who T-Shirt", Size: product.L}, Price: 29.95},
			}

			expected := []Inventory{
				{Product: product.Product{Name: "Dr. Who T-Shirt", Size: product.L}, Price: 29.95},
				{Product: product.Product{Name: "Widget 1", Size: product.NA}, Price: 199.99},
				{Product: product.Product{Name: "Widget 2", Size: product.NA}, Price: 59.75},
			}

			sort.Sort(InventoryComparator(actual))
			So(actual, ShouldResemble, expected)
		})

		Convey("Sorts by size when the names are the same", func() {
			actual := []Inventory{
				{Product: product.Product{Name: "Widget 1", Size: product.NA}, Price: 59.75},
				{Product: product.Product{Name: "Dr. Who T-Shirt", Size: product.L}, Price: 29.95},
				{Product: product.Product{Name: "Dr. Who T-Shirt", Size: product.S}, Price: 29.95},
				{Product: product.Product{Name: "Dr. Who T-Shirt", Size: product.M}, Price: 29.95},
			}

			expected := []Inventory{
				{Product: product.Product{Name: "Dr. Who T-Shirt", Size: product.S}, Price: 29.95},
				{Product: product.Product{Name: "Dr. Who T-Shirt", Size: product.M}, Price: 29.95},
				{Product: product.Product{Name: "Dr. Who T-Shirt", Size: product.L}, Price: 29.95},
				{Product: product.Product{Name: "Widget 1", Size: product.NA}, Price: 59.75},
			}

			sort.Sort(InventoryComparator(actual))
			So(actual, ShouldResemble, expected)
		})
	})
}
