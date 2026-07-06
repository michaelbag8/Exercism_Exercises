package main

import "fmt"

func FavoriteCards() []int {
	var result []int
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range nums {
		if v == 2 || v == 6 || v == 9 {
			result = append(result, v)
		}
	}
	return result
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	returnItem := 0
	if index < 0 || index > len(slice) {
		return -1
	}
	for i, v := range slice {
		if i == index {
			returnItem = v
		}
	}
	return returnItem

}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < 0 || index >= len(slice){
		slice = append(slice, index)
	}

	for i := range slice {
		if i == index {
			slice[i] = value

		}

	}

	return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	var result []int
	result = append(result, values...)
	slice = append(result, slice...)
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	for i := range slice {
		if i == index {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func main() {
	fmt.Println(FavoriteCards())
	fmt.Println("----------------")
	card := GetItem([]int{1, 2, 4, 1}, -5)
	fmt.Println(card) //4-----2, 3-----0, 2-----1,1-----3, 0----0
	fmt.Println("----------------")

	index := 2
	newCard := 6
	cards := SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards)

	fmt.Println("----------------")
	slice := []int{3, 2, 6, 4, 8}
	cardz := PrependItems(slice, 5, 1)
	fmt.Println(cardz)
	// Output: [5 1 3 2 6 4 8]

	fmt.Println("----------------")
	cardr := RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println(cardr)
	// Output: [3 2 4 8]

	fmt.Println("----------------")
	sl := []int{5, 2, 10, 6, 8, 7, 0, 9}
	i := 8
	value := 8
	fmt.Println(SetItem(sl, i, value))
	//{5, 2, 10, 6, 8, 7, 0, 9, 8}
}

/*

package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if index >= len(slice) || index < 0{
        return -1
    } 
    return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index >= len(slice) || index < 0{
        return append(slice, value)
    }
	slice[index] = value
    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if index >= len(slice) || index < 0 {
        return slice
    }
	return append(slice[:index], slice[index+1:]...)
}

*/