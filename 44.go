package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "reza"

	fmt.Println(name)

	fmt.Println(reflect.TypeOf(name))
}
