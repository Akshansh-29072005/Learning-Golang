package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int {2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice){
        return -1
    } else {
        return slice[index]
    }
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(stack []int, index int, newCard int) []int {
	if index < 0 || index >= len(stack) {
		// Index out of bounds, append the new card
		return append(stack, newCard)
	}
	// Index is within bounds, set the item at that position
	stack[index] = newCard
	return stack
}

func PrependItems(stack []int, values ...int) []int {
	// 'values' is treated as a slice []int inside the function.
	// We append the original stack to the end of the new items slice.
	// Note the '...' operator is used here to unpack the 'stack' slice into individual arguments for append.
	return append(values, stack...)
}

// RemoveItem removes the card at position index from the stack and returns the stack.
// If the index is out of bounds, the stack is returned unchanged.
func RemoveItem(stack []int, index int) []int {
	if index < 0 || index >= len(stack) {
		// Index out of bounds, return the original stack unchanged
		return stack
	}
	return append(stack[:index], stack[index+1:]...)
}
