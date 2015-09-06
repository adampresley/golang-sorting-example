package inventory

type InventoryComparator []Inventory

func (items InventoryComparator) Len() int {
	return len(items)
}

func (items InventoryComparator) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items InventoryComparator) Less(i, j int) bool {
	if items[i].Name != items[j].Name {
		return items[i].Name < items[j].Name
	} else {
		return items[i].Size.SortOrder < items[j].Size.SortOrder
	}
}
