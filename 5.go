package main

import (
	"fmt"
	"reflect"
)

func main() {

	number := 12.5

	fmt.Println(number)

	fmt.Println(reflect.TypeOf(number))
}
