package sandboxgo

import (
	"strconv"
)

func Checkout(name string, number_of_room int) string {
	return order(name, number_of_room)
}

func order(nama string, number_of_room int) string {
	return nama + " order " + strconv.Itoa(number_of_room) + " room with price " + strconv.Itoa(price(number_of_room))
}

func price(number_of_room int) int {
	price := number_of_room * 350000
	return price
}
