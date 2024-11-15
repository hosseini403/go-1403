package main

import (
	"fmt"
	"reflect"
)

func main() {

	number := 13.5

	fmt.Println(number)

	fmt.Println(reflect.TypeOf(number))
}
