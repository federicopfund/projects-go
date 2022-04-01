package controller



func Append(items []int, i int) []int {
	if len(items) == cap(items) { // No more space in underlying array
		// Go has a better growth heuristic than adding 1 every append
		newItems := make([]int, len(items)+1)
		copy(newItems, items)
		items = newItems
	} else {
		items = items[:len(items)+1]
	}

	items[len(items)-1] = i
	return items
}