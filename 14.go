package main

import (
	"fmt"
	"reflect"
)

func main() {

	A := 7
	B := 5

	fmt.Println(A != B)

	z := true

	fmt.Println(!z)

	x := 1
	y := 2

	fmt.Println((x < y) && (true))

	my_float_number := 8.5
	fmt.Println("my float number:", my_float_number)
	my_int_number := int(my_float_number)
	fmt.Println("my int number:", my_int_number)
	fmt.Println(reflect.TypeOf(my_int_number))
}
