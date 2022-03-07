package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {

	if index < 0 || index >= len(slice) {

		return 0, false
	}
	return slice[index], true
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	card, isValid := GetItem(slice, index)
	if card == 0 && !isValid {
		updatedSlice := append(slice, value)
		return updatedSlice

	}

	slice[index] = value
	return slice

}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	if length <= 0 {
		emptySlice := []int{}
		return emptySlice
	}
	slice := make([]int, length)

	for i := 0; i < length; i++ {
		slice[i] = value
	}

	return slice

}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	card, isValid := GetItem(slice, index)
	if card == 0 && !isValid {
		return slice

	}
	updatedSlice := append(slice[:index], slice[index+1:]...)
	return updatedSlice

}
