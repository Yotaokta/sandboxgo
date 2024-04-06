package sandboxgo

import (
	"strconv"
)

func Checkout(name string, number_of_room int) string {
	return order(name, number_of_room)
}

func order(nama string, number_of_room int) string {
	return nama + " order " + strconv.Itoa(number_of_room) + " room"
}
