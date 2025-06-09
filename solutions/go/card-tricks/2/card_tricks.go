package cards

func FavoriteCards() []int {
	return []int{2,6,9}
}

func GetItem(slice []int, index int) int {
	if len(slice) <= index || index < 0 {
        return -1
    }
    return slice[index]
}

func SetItem(slice []int, index, value int) []int {
	if len(slice) <= index || index < 0 {
        slice = append(slice, value)
        return slice
    }
    slice[index] = value
    return slice
}

func PrependItems(slice []int, values ...int) []int {
	slice = append(values, slice...)
    return slice
}

func RemoveItem(slice []int, index int) []int {
	if len(slice) <= index || index < 0 {
        return slice
    }
    slice = append(slice[:index], slice[index+1:]...)
    return slice
}
