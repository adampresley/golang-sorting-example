package product

type ProductSize struct {
	Abbreviation string
	Size         string
	SortOrder    int
}

var NA ProductSize = ProductSize{"NA", "Not applicable", 1}
var XS ProductSize = ProductSize{"XS", "Extra Small", 2}
var S ProductSize = ProductSize{"S", "Small", 3}
var M ProductSize = ProductSize{"M", "Medium", 4}
var L ProductSize = ProductSize{"L", "Large", 5}
var XL ProductSize = ProductSize{"XL", "Extra Large", 6}
var XXL ProductSize = ProductSize{"XXL", "2-Extra Large", 7}
