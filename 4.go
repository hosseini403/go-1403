package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Ali"

	fmt.Println(name)

	fmt.Println(reflect.TypeOf(name))
}
