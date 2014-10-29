package main

import (
	"fmt"
	"reflect"
)

func main() {
	a()
}
func a() {
	d := reflect.Func()
	fmt.Println(d)
}
