package product

type ProductSizeComparator []ProductSize

func (items ProductSizeComparator) Len() int {
	return len(items)
}

func (items ProductSizeComparator) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items ProductSizeComparator) Less(i, j int) bool {
	return items[i].SortOrder < items[j].SortOrder
}
